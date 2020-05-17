package svc

import (
	"errors"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
	"github.com/sirupsen/logrus"
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
	s.in++

	var sid = fmt.Sprintf("%v", stream.Context().Value("AccountID"))
	var buf = make(chan *pb.StreamChatEvent, 1000)
	var err error
	var name = "Replace"

	s.withReadLock(func() {
		var ok bool
		_, ok = s.profile[sid]
		if !ok {
			err = errors.New("not authorized")
			return
		}
		if _, ok := s.buf[sid]; !ok {
			err = errors.New("not authorized")
			return
		}
	})
	if err != nil {
		return err
	}

	go s.withReadLock(func() {
		// logrus.Debugf("Log message=%s", in.Payload.Text)
		for _, buf := range s.buf {
			buf <- &pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_Message{
					Message: &pb.EventMessage{
						// Payload: out.Result,
					},
				},
			}
		}
	})

	go func() {
		time.Sleep(5 * time.Second)
		s.withWriteLock(func() {
			if _, ok := s.buf[sid]; ok {
				return
			}
			s.unsafeExpire(sid)
		})
	}()

	s.withWriteLock(func() {
		var ok bool
		_, ok = s.profile[sid]
		if !ok {
			err = errors.New("not authorized")
			return
		}
		if _, ok := s.buf[sid]; ok {
			err = errors.New("already connected")
			return
		}
		s.buf[sid] = buf
	})

	if err != nil {
		return err
	}

	go s.withReadLock(func() {
		logrus.Debugf("Join name=%s", name)
		for _, buf := range s.buf {
			buf <- &pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_Join{
					Join: &pb.EventJoin{},
				},
			}
		}
	})

	defer s.withReadLock(func() {
		logrus.Debugf("Leave name=%s", name)
		for _, buf := range s.buf {
			buf <- &pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_Leave{
					Leave: &pb.EventLeave{},
				},
			}
		}
	})

	defer s.withWriteLock(func() { s.unsafeExpire(sid) })

	tick := time.Tick(time.Second)

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case event := <-buf:
			if err := stream.Send(event); err != nil {
				return err
			}
			s.out++
		case <-tick:
			if err := stream.Send(
				&pb.StreamChatEvent{
					Event: &pb.StreamChatEvent_None{
						None: &pb.EventNone{},
					},
				}); err != nil {
				return err
			}
			s.out++
		}
	}
}

func (s *chatServer) withReadLock(f func()) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f()
}

func (s *chatServer) withWriteLock(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}

func (s *chatServer) unsafeExpire(sid string) {
	if buf, ok := s.buf[sid]; ok {
		close(buf)
	}
	delete(s.profile, sid)
	delete(s.buf, sid)
}
