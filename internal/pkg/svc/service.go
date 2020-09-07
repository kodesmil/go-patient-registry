package svc

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/infobloxopen/protoc-gen-gorm/types"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
	"github.com/kodesmil/ks-backend/internal/pkg/strings"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
)

func NewServicesServer(database *gorm.DB) (pb.ServicesServer, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	m, err := app.Messaging(context.Background())
	if err != nil {
		log.Errorf("error getting messaging client: %v", err)
	}
	return &servicesServer{
		ServicesDefaultServer: &pb.ServicesDefaultServer{DB: database},
		connections:           make(map[string]*ServiceStreamConnection),
		fireApp:               app,
		fireMessaging:         m,
	}, nil
}

func (s *servicesServer) ListServiceSession(ctx context.Context, in *pb.ListServiceSessionRequest) (*pb.ListServiceSessionResponse, error) {
	return s.CustomListServiceSession(ctx, in)
}

func (s *servicesServer) ListServiceOfferSession(ctx context.Context, in *pb.ListServiceOfferSessionRequest) (*pb.ListServiceOfferSessionResponse, error) {
	return s.CustomListServiceOfferSession(ctx, in)
}

type servicesServer struct {
	*pb.ServicesDefaultServer
	connections   map[string]*ServiceStreamConnection
	fireApp       *firebase.App
	fireMessaging *messaging.Client
}

type ServiceStreamConnection struct {
	stream pb.Services_BiDiServer
	active bool
	error  chan error
}

func (s *servicesServer) BroadcastMessage(ids []string, msg *pb.StreamSessionOutputEvent) error {
	wait := sync.WaitGroup{}
	for id, conn := range s.connections {
		if !strings.Include(ids, id) {
			continue
		}
		wait.Add(1)
		go func(msg *pb.StreamSessionOutputEvent, conn *ServiceStreamConnection) {
			defer wait.Done()
			if conn.active {
				err := conn.stream.Send(msg)
				if err != nil {
					log.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}
	return nil
}

func (s *servicesServer) BiDi(stream pb.Services_BiDiServer) error {

	accountID := fmt.Sprintf(
		"%v",
		stream.Context().Value("AccountID"),
	)

	if res, ok := s.connections[accountID]; !ok || !res.active {
		conn := &ServiceStreamConnection{
			stream: stream,
			active: true,
			error:  make(chan error),
		}
		s.connections[accountID] = conn
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.GetRequestSession() != nil {
			msg := in.GetRequestSession()
			session := msg.Session

			profiles := []string{
				session.ProfileId,
				session.Offer.Provider.Employments[0].ProfileId,
			}

			if err := s.BroadcastMessage(profiles, &pb.StreamSessionOutputEvent{
				Event: &pb.StreamSessionOutputEvent_SessionRequested{
					SessionRequested: &pb.EventSessionRequested{
						Session: session,
					},
				},
			}); err != nil {
				return err
			}

			if session.Status == pb.ServiceSession_NOT_STARTED {
				if err := s.SendNotification(notification{
					title:     "New session request!",
					body:      fmt.Sprintf("%s asked you for an advise on %s", session.ProfileId, session.Offer.Title),
					profileId: session.Offer.Provider.Employments[0].ProfileId,
					sessionId: session.Id.Value,
				}); err != nil {
					return err
				}

				if err := s.SendNotification(notification{
					title:     fmt.Sprintf("Connecting with %s", session.Offer.Provider.Details.Name),
					body:      "Waiting for an answer...",
					profileId: session.ProfileId,
					sessionId: session.Id.Value,
				}); err != nil {
					return err
				}
			}

			if session.Status == pb.ServiceSession_NOT_STARTED && accountID == session.Offer.Owner.ProfileId {
				session.Status = pb.ServiceSession_ONGOING
				session, err = pb.DefaultStrictUpdateServiceSession(
					stream.Context(), session, s.DB,
				)
				if err := s.SendNotification(notification{
					title:     "Connected",
					body:      "XXX has been successfully connected",
					profileId: session.Offer.Owner.ProfileId,
					sessionId: session.Id.Value,
				}); err != nil {
					return err
				}

				if err := s.SendNotification(notification{
					title:     "Connected",
					body:      "XXX has been successfully connected",
					profileId: session.ProfileId,
					sessionId: session.Id.Value,
				}); err != nil {
					return err
				}
			}

			go func(sessionID *types.UUIDValue) {

				defer func() {
					session, err = pb.DefaultReadServiceSession(
						context.Background(),
						&pb.ServiceSession{Id: sessionID},
						s.DB,
					)
					if session.Status == pb.ServiceSession_NOT_STARTED {
						session.Status = pb.ServiceSession_CANCELED
						session.FinishedAt = &timestamp.Timestamp{Seconds: time.Now().Unix()}
						session, err = pb.DefaultStrictUpdateServiceSession(
							context.Background(), session, s.DB,
						)
						if err != nil {
							log.Error(err)
						}
					}
				}()

				for {
					session, err = pb.DefaultReadServiceSession(
						context.Background(),
						&pb.ServiceSession{Id: sessionID},
						s.DB,
					)
					if err != nil || stream.Context().Err() != nil {
						log.Error(err)
						return
					}
					if session.Status != pb.ServiceSession_NOT_STARTED {
						log.Infof("Session changed status: %v", session.Status)
						return
					}
					log.Infof("Waiting to connect service session...")
					if err := s.SendNotification(notification{
						title:     fmt.Sprintf("Connecting with %s", session.Offer.Provider.Details.Name),
						body:      "Waiting for an answer...",
						profileId: session.Offer.Owner.ProfileId,
						sessionId: session.Id.Value,
					}); err != nil {
						return
					}
					time.Sleep(time.Second * 5)
				}
			}(session.Id)
		}
	}
}

func (s *servicesServer) SendNotification(n notification) error {
	go func() {
		var notificationDevices []pb.NotificationDeviceORM
		if err := s.DB.
			Select("notification_devices.device_token").
			Joins("left join notification_settings ns on ns.profile_id = notification_devices.profile_id").
			Where("ns.enable_notifications = ?", true).
			Where("ns.profile_id = ?", n.profileId).
			Find(&notificationDevices).
			Error; err != nil {
			log.Fatal(err)
		}

		if len(notificationDevices) == 0 {
			log.Info("no notification devices are registered on user device")
			return
		}
		ttl := time.Second * 15
		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: n.title,
				Body:  n.body,
			}, Android: &messaging.AndroidConfig{
				TTL: &ttl,
				Data: map[string]string{
					"route": fmt.Sprintf("/sessions/%s", n.sessionId),
				},
				Notification: &messaging.AndroidNotification{
					Tag:         "service_stream",
					ClickAction: "FLUTTER_NOTIFICATION_CLICK",
				},
			},
			Tokens: pb.MapToDeviceTokens(notificationDevices),
		}
		response, err := s.fireMessaging.SendMulticast(context.Background(), message)
		if err != nil {
			log.Error(err)
			return
		}
		log.Println("Successfully sent message:", response)
	}()
	return nil
}

type notification struct {
	title     string
	body      string
	profileId string
	sessionId string
}

func MapProfiles(vs []*pb.Profile, f func(*pb.Profile) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
