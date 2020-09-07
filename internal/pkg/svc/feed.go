package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
)

func NewFeedServer(database *gorm.DB) (pb.FeedServer, error) {
	return &feedServer{&pb.FeedDefaultServer{DB: database}}, nil
}

type feedServer struct {
	*pb.FeedDefaultServer
}
