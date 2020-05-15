package svc

import (
	"context"
	"fmt"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
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
