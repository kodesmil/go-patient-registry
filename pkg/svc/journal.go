package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewJournalEntriesServer(database *gorm.DB) (pb.JournalEntriesServer, error) {
	return &journalEntriesServer{&pb.JournalEntriesDefaultServer{DB: database}}, nil
}

type journalEntriesServer struct {
	*pb.JournalEntriesDefaultServer
}

func NewJournalSubjectsServer(database *gorm.DB) (pb.JournalSubjectsServer, error) {
	return &journalSubjectsServer{&pb.JournalSubjectsDefaultServer{DB: database}}, nil
}

type journalSubjectsServer struct {
	*pb.JournalSubjectsDefaultServer
}
