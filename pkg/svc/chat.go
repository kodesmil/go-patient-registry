package svc

import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
	"github.com/kodesmil/ks-backend/pkg/strings"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
)

func NewChatServer(database *gorm.DB) *chatServer {
	return &chatServer{
		database:    database,
		connections: make(map[string]*Connection),
	}
}

type chatServer struct {
	database    *gorm.DB
	connections map[string]*Connection
}

type Connection struct {
	stream pb.Chat_StreamServer
	active bool
	error  chan error
}

func (s *chatServer) BroadcastMessage(ids []string, msg *pb.StreamChatEvent) error {
	wait := sync.WaitGroup{}
	for id, conn := range s.connections {
		if !strings.Include(ids, id) {
			continue
		}
		wait.Add(1)
		go func(msg *pb.StreamChatEvent, conn *Connection) {
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

func (s *chatServer) Stream(stream pb.Chat_StreamServer) error {

	accountID := &resource.Identifier{ResourceId: fmt.Sprintf(
		"%v",
		stream.Context().Value("AccountID"),
	)}

	if res, ok := s.connections[accountID.ResourceId]; !ok || !res.active {
		conn := &Connection{
			stream: stream,
			active: true,
			error:  make(chan error),
		}
		s.connections[accountID.ResourceId] = conn
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.GetLoadRooms() != nil {
			filtering := &query.Filtering{
				Root: &query.Filtering_StringCondition{
					StringCondition: &query.StringCondition{
						FieldPath: []string{"chat_room_participants", "profile_id"},
						Value:     accountID.ResourceId,
						Type:      query.StringCondition_EQ,
					},
				},
			}
			db := s.database.Joins("left join chat_room_participants on chat_room_participants.chat_room_id::uuid = chat_rooms.id")
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
			db := s.database.
				Joins("left join chat_room_participants on chat_room_participants.id = chat_messages.author_id").
				Joins("left join chat_rooms on chat_rooms.id = chat_room_participants.chat_room_id::uuid")
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
						Messages: out,
					},
				},
			}); err != nil {
				return err
			}
			err = db.Exec(
				"UPDATE chat_room_participants SET last_seen_at=? WHERE id=?",
				time.Now(),
				room.Me.Id.ResourceId,
			).Error
			if err != nil {
				return err
			}
		}
		if in.GetSendMessage() != nil {
			message := in.GetSendMessage().Message
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

			var profiles []profileId
			err = s.database.
				Table("chat_room_participants").
				Select("profile_id").
				Where("chat_room_id = ?", out.GetAuthor().ChatRoom.Id.ResourceId).
				Find(&profiles).
				Error

			if err != nil {
				return err
			}

			if err := s.BroadcastMessage(Map(profiles, func(profile profileId) string {
				return profile.ProfileId
			}), &pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_SendMessage{
					SendMessage: &pb.EventSendMessage{
						Message: out,
					},
				},
			}); err != nil {
				return err
			}
		}
	}
}

type profileId struct {
	ProfileId string
}

func Map(vs []profileId, f func(profileId) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
