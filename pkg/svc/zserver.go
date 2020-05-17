package svc

import (
	"errors"
	"flag"
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
	"github.com/sirupsen/logrus"
	"math/rand"
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

// NewChatMessagesServer returns an instance of the default feed articles server interface

func NewChatMessagesServer(database *gorm.DB) (pb.ChatMessagesServer, error) {
	return &chatMessagesServer{
		ChatMessagesDefaultServer: &pb.ChatMessagesDefaultServer{DB: database},
		profile:                   make(map[string]pb.Profile),
		buf:                       make(map[string]chan *pb.StreamChatEvent),

		last: time.Now(),
	}, nil
}

type chatMessagesServer struct {
	*pb.ChatMessagesDefaultServer

	mu      sync.RWMutex
	profile map[string]pb.Profile
	buf     map[string]chan *pb.StreamChatEvent

	last time.Time
	in   int64
	out  int64
}

func (s *chatMessagesServer) Stream(ctx context.Context, in *pb.StreamConnectRequest) (*pb.StreamChatEvent, error) {
	s.in++

	var (
		buf  = make(chan *pb.StreamChatEvent, 1000)
		sid  = fmt.Sprintf("%v", ctx.Value("AccountID"))
		err  error
		name string
	)

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
		return nil, err
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
			if err := stream.Send(&pb.Event{None: &pb.EventNone{}}); err != nil {
				return err
			}
			s.out++
		}
	}

	out, err := s.ChatMessagesDefaultServer.Stream(ctx, in)

	return out, err
}

func (s *chatMessagesServer) Create(ctx context.Context, in *pb.CreateChatMessageRequest) (*pb.CreateChatMessageResponse, error) {
	s.in++

	accountId := fmt.Sprintf("%v", ctx.Value("AccountID"))
	var err error

	s.withReadLock(func() {
		var ok bool
		_, ok = s.profile[accountId]
		if !ok {
			err = errors.New("not authorized")
			return
		}
		if _, ok := s.buf[accountId]; !ok {
			err = errors.New("not authorized")
			return
		}
	})
	if err != nil {
		return nil, err
	}

	if len(in.Payload.Text) == 0 {
		return nil, errors.New("message must be not empty")
	}
	if len(in.Payload.Text) > 140 {
		return nil, errors.New("message must be less than or equal 140 characters")
	}

	out, err := s.ChatMessagesDefaultServer.Create(ctx, in)

	go s.withReadLock(func() {
		logrus.Debugf("Log message=%s", in.Payload.Text)
		for _, buf := range s.buf {
			buf <- &pb.StreamChatEvent{
				Event: &pb.StreamChatEvent_Log{
					Log: &pb.EventLog{
						Payload: out.Result,
					},
				},
			}
		}
	})

	return out, err
}

func (s *chatMessagesServer) withReadLock(f func()) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	f()
}

func (s *chatMessagesServer) withWriteLock(f func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f()
}

func (s *chatMessagesServer) unsafeExpire(sid string) {
	if buf, ok := s.buf[sid]; ok {
		close(buf)
	}
	delete(s.profile, sid)
	delete(s.buf, sid)
}
