package svc

import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/query"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
	"github.com/kodesmil/ks-backend/internal/pkg/strings"
	log "github.com/sirupsen/logrus"
	"io"
	"sync"
	"time"
)

func NewChatServer(database *gorm.DB) (*chatServer, error) {
	return &chatServer{
		ChatDefaultServer: &pb.ChatDefaultServer{DB: database},
		connections:       make(map[string]*Connection),
	}, nil
}

type chatServer struct {
	*pb.ChatDefaultServer
	connections map[string]*Connection
}

type Connection struct {
	stream pb.Chat_BiDiServer
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

func (s *chatServer) BiDi(stream pb.Chat_BiDiServer) error {
	accountID := fmt.Sprintf(
		"%v",
		stream.Context().Value("AccountID"),
	)

	if res, ok := s.connections[accountID]; !ok || !res.active {
		conn := &Connection{
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
		if in.GetLoadRooms() != nil {
			db := s.DB.Joins("left join chat_room_participants on chat_room_participants.chat_room_id = chat_rooms.id")
			out, err := pb.DefaultListChatRoom(stream.Context(), db, &query.Filtering{
				Root: &query.Filtering_StringCondition{
					StringCondition: &query.StringCondition{
						FieldPath: []string{"chat_room_participants", "profile_id"},
						Value:     accountID,
						Type:      query.StringCondition_EQ,
					},
				},
			}, nil, nil, nil)
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
						Value:     room.Room.Id.Value,
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
			db := s.DB.
				Joins("left join chat_room_participants on chat_room_participants.id = chat_.author_id").
				Joins("left join chat_rooms on chat_rooms.id = chat_room_participants.chat_room_id")
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
				room.Me.Id,
			).Error
			if err != nil {
				return err
			}
		}
		if in.GetSendMessage() != nil {
			message := in.GetSendMessage().Message
			var created, err = pb.DefaultCreateChatMessage(
				stream.Context(), message, s.DB,
			)
			if err != nil {
				return err
			}
			out, err := pb.DefaultReadChatMessage(
				stream.Context(), &pb.ChatMessage{Id: created.Id},
				s.DB,
			)
			if err != nil {
				return err
			}

			var profiles []profileId
			err = s.DB.
				Table("chat_room_participants").
				Select("profile_id").
				Where("chat_room_id = ?", out.GetAuthor().ChatRoom.Id.Value).
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
