package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewNotificationDevicesServer(database *gorm.DB) (pb.NotificationDevicesServer, error) {
	return &notificationDevicesServer{&pb.NotificationDevicesDefaultServer{DB: database}}, nil
}

type notificationDevicesServer struct {
	*pb.NotificationDevicesDefaultServer
}

func NewNotificationSettingsServer(database *gorm.DB) (pb.NotificationSettingsServer, error) {
	return &notificationSettingsServer{&pb.NotificationSettingsDefaultServer{DB: database}}, nil
}

type notificationSettingsServer struct {
	*pb.NotificationSettingsDefaultServer
}
