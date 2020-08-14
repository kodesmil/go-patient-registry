package svc

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/infobloxopen/protoc-gen-gorm/types"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
	"github.com/kodesmil/ks-backend/internal/pkg/strings"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
)

func NewServiceSessionStreamServer(database *gorm.DB) *serviceSessionStreamServer {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	messaging, err := app.Messaging(context.Background())
	if err != nil {
		log.Errorf("error getting messaging client: %v", err)
	}
	return &serviceSessionStreamServer{
		database:      database,
		connections:   make(map[string]*ServiceStreamConnection),
		fireApp:       app,
		fireMessaging: messaging,
	}
}

type serviceSessionStreamServer struct {
	database      *gorm.DB
	connections   map[string]*ServiceStreamConnection
	fireApp       *firebase.App
	fireMessaging *messaging.Client
}

type ServiceStreamConnection struct {
	stream pb.ServiceSessionStream_BiDiServer
	active bool
	error  chan error
}

func (s *serviceSessionStreamServer) BroadcastMessage(ids []string, msg *pb.StreamSessionOutputEvent) error {
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

func (s *serviceSessionStreamServer) BiDi(stream pb.ServiceSessionStream_BiDiServer) error {

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

			if err := s.SendNotification(notification{
				title:     "You received new session request!",
				body:      "XXX asked you for an advise on...",
				profileId: session.Offer.Provider.Employments[0].ProfileId,
			}); err != nil {
				return err
			}

			if err := s.SendNotification(notification{
				title:     "Session request sent to XXX",
				body:      "Waiting for getting connected...",
				profileId: session.ProfileId,
			}); err != nil {
				return err
			}
			go func(sessionID *types.UUIDValue) {

				session, err = pb.DefaultReadServiceSession(
					context.Background(),
					&pb.ServiceSession{Id: sessionID},
					s.database,
				)

				defer func() {
					if session.Status == pb.ServiceSession_NOT_STARTED {
						session.Status = pb.ServiceSession_FINISHED // cancelled
						session, err = pb.DefaultStrictUpdateServiceSession(context.Background(), session, s.database)
						if err != nil {
							log.Error(err)
						}
					}
				}()

				for {
					if err != nil || stream.Context().Err() != nil {
						return
					}
					if session.Status != pb.ServiceSession_NOT_STARTED {
						log.Infof("Session changed status: %v", session.Status)
						return
					}
					log.Infof("Waiting to connect service session...")
					if err := s.SendNotification(notification{
						title:     "Waiting for connection",
						body:      fmt.Sprintf("%d seconds passed", time.Now().Unix()-session.CreatedAt.Seconds),
						profileId: session.Offer.Provider.Employments[0].ProfileId,
					}); err != nil {
						return
					}
					time.Sleep(time.Second * 5)
				}
			}(session.Id)
		} else if in.GetJoinSession() != nil {
			msg := in.GetJoinSession()
			session := msg.Session
			session.Status = pb.ServiceSession_ONGOING
			session, err = pb.DefaultStrictUpdateServiceSession(stream.Context(), session, s.database)
			if err := s.SendNotification(notification{
				title:     "Connected",
				body:      "XXX has been successfully connected",
				profileId: session.Offer.Provider.Employments[0].ProfileId,
			}); err != nil {
				return err
			}

			if err := s.SendNotification(notification{
				title:     "Connected",
				body:      "XXX has been successfully connected",
				profileId: session.ProfileId,
			}); err != nil {
				return err
			}
		}
	}
}

func (s *serviceSessionStreamServer) SendNotification(n notification) error {
	go func() {
		var notificationDevices []pb.NotificationDeviceORM
		if err := s.database.
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
		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: n.title,
				Body:  n.body,
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
}

func MapProfiles(vs []*pb.Profile, f func(*pb.Profile) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
