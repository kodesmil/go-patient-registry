package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
)

// NewProfilesServer returns an instance of the default profiles server interface
func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
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
