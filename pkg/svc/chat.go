package svc

import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
	"github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
)

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
			room := in.GetLoadRoom()
			filtering := &query.Filtering{
				Root: &query.Filtering_StringCondition{
					StringCondition: &query.StringCondition{
						FieldPath: []string{"chat_rooms", "id"},
						Value:     room.Room.Id.ResourceId,
						Type:      query.StringCondition_EQ,
					},
				},
			}
			sorting := &query.Sorting{
				Criterias: []*query.SortCriteria{
					{
						Tag:   "created_at",
						Order: query.SortCriteria_DESC,
					},
				},
			}
			db := s.database.Joins("left join chat_rooms on chat_rooms.id = chat_messages.chat_room_id::uuid")
			out, err := pb.DefaultListChatMessage(
				stream.Context(), db,
				filtering, sorting, nil, nil,
			)
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
			var created, err = pb.DefaultCreateChatMessage(
				stream.Context(), message, s.database,
			)
			if err != nil {
				return err
			}
			out, err := pb.DefaultReadChatMessage(
				stream.Context(), &pb.ChatMessage{Id: created.Id},
				s.database,
			)
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
}
