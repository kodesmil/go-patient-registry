package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
)

func NewJournalServer(database *gorm.DB) (pb.JournalServer, error) {
	return &journalServer{&pb.JournalDefaultServer{DB: database}}, nil
}

type journalServer struct {
	*pb.JournalDefaultServer
}
