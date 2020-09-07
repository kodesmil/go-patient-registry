package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
)

func NewNotificationsServer(database *gorm.DB) (pb.NotificationsServer, error) {
	return &notificationsServer{&pb.NotificationsDefaultServer{DB: database}}, nil
}

type notificationsServer struct {
	*pb.NotificationsDefaultServer
}
