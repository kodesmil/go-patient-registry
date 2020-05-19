package svc

import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"

	"golang.org/x/net/context"
)

// NewProfilesServer returns an instance of the default profiles server interface

func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
}

func (s *profilesServer) Create(ctx context.Context, in *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	in.Payload.Id = &resource.Identifier{ResourceId: fmt.Sprintf("%v", ctx.Value("AccountID"))}
	return s.ProfilesDefaultServer.Create(ctx, in)
}

type profilesServer struct {
	*pb.ProfilesDefaultServer
}

// NewGroupsServer returns an instance of the default groups server interface

func NewGroupsServer(database *gorm.DB) (pb.GroupsServer, error) {
	return &groupsServer{&pb.GroupsDefaultServer{DB: database}}, nil
}

type groupsServer struct {
	*pb.GroupsDefaultServer
}

// NewFeedArticlesServer returns an instance of the default feed articles server interface

func NewFeedArticlesServer(database *gorm.DB) (pb.FeedArticlesServer, error) {
	return &feedArticlesServer{&pb.FeedArticlesDefaultServer{DB: database}}, nil
}

type feedArticlesServer struct {
	*pb.FeedArticlesDefaultServer
}

// NewFeedArticleDetailsServer returns an instance of the default feed articles server interface

func NewFeedArticleDetailsServer(database *gorm.DB) (pb.FeedArticleDetailsServer, error) {
	return &feedArticleDetailsServer{&pb.FeedArticleDetailsDefaultServer{DB: database}}, nil
}

type feedArticleDetailsServer struct {
	*pb.FeedArticleDetailsDefaultServer
}

// NewJournalEntriesServer returns an instance of the default journal server interface

func NewJournalEntriesServer(database *gorm.DB) (pb.JournalEntriesServer, error) {
	return &journalEntriesServer{&pb.JournalEntriesDefaultServer{DB: database}}, nil
}

type journalEntriesServer struct {
	*pb.JournalEntriesDefaultServer
}

// NewJournalSubjectsServer returns an instance of the default journal server interface

func NewJournalSubjectsServer(database *gorm.DB) (pb.JournalSubjectsServer, error) {
	return &journalSubjectsServer{&pb.JournalSubjectsDefaultServer{DB: database}}, nil
}

type journalSubjectsServer struct {
	*pb.JournalSubjectsDefaultServer
}

// NewNotificationDevicesServer returns an instance of the default feed articles server interface

func NewNotificationDevicesServer(database *gorm.DB) (pb.NotificationDevicesServer, error) {
	return &notificationDevicesServer{&pb.NotificationDevicesDefaultServer{DB: database}}, nil
}

type notificationDevicesServer struct {
	*pb.NotificationDevicesDefaultServer
}

// NewNotificationSettingsServer returns an instance of the default feed articles server interface

func NewNotificationSettingsServer(database *gorm.DB) (pb.NotificationSettingsServer, error) {
	return &notificationSettingsServer{&pb.NotificationSettingsDefaultServer{DB: database}}, nil
}

type notificationSettingsServer struct {
	*pb.NotificationSettingsDefaultServer
}

// NewChatServer returns an instance of the default feed articles server interface

func NewChatServer(database *gorm.DB) *chatServer {
	return &chatServer{
		database: database,
		profile:  make(map[string]pb.Profile),
		buf:      make(map[string]chan *pb.StreamChatEvent),
		last:     time.Now(),
	}
}

type chatServer struct {
	database *gorm.DB

	mu      sync.RWMutex
	profile map[string]pb.Profile
	buf     map[string]chan *pb.StreamChatEvent

	last time.Time
	in   int64
	out  int64
}

/*

// RouteChat receives a stream of message/location pairs, and responds with a stream of all
// previous messages at each of those locations.
func (s *routeGuideServer) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serialize(in.Location)

		s.mu.Lock()
		s.routeNotes[key] = append(s.routeNotes[key], in)
		// Note: this copy prevents blocking other clients while serving this one.
		// We don't need to do a deep copy, because elements in the slice are
		// insert-only and never modified.
		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

*/

func (s *chatServer) Stream(stream pb.Chat_StreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		accountID := &resource.Identifier{ResourceId: fmt.Sprintf(
			"%v",
			stream.Context().Value("AccountID"),
		)}
		logrus.Info(accountID)
		if in.GetLoadRooms() != nil {
			filtering := &query.Filtering{
				Root: &query.Filtering_StringCondition{
					StringCondition: &query.StringCondition{
						FieldPath: []string{"chat_room_profiles", "profile_id"},
						Value:     accountID.ResourceId,
						Type:      query.StringCondition_EQ,
					},
				},
			}
			db := s.database.Joins("left join chat_room_profiles on chat_room_profiles.chat_room_id = chat_rooms.id")
			out, err := pb.DefaultListChatRoom(stream.Context(), db, filtering, nil, nil, nil)
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_SendRooms{
					SendRooms: &pb.EventSendRooms{
						Rooms: out,
					},
				},
			}); err != nil {
				return err
			}
		}
		if in.GetLoadRoom() != nil {
			out, err := pb.DefaultListChatMessage(stream.Context(), s.database, nil, nil, nil, nil)
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_SendMessages{
					SendMessages: &pb.EventSendMessages{
						Payload: out,
					},
				},
			}); err != nil {
				return err
			}
		}
		if in.GetSendMessage() != nil {
			message := in.GetSendMessage().Payload
			out, err := pb.DefaultCreateChatMessage(stream.Context(), message, s.database)
			if err != nil {
				return err
			}
			if err := stream.Send(&pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_SendMessage{
					SendMessage: &pb.EventSendMessage{
						Payload: out,
					},
				},
			}); err != nil {
				return err
			}
		}
	}

	for _, buf := range s.buf {
		buf <- &pb.StreamChatEvent{
			Event: &pb.StreamChatEvent_LeaveRooms{
				LeaveRooms: &pb.EventLeaveRooms{},
			},
		}
	}

	return nil
}
