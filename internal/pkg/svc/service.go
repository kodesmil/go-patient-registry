package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
)

func NewServicesServer(database *gorm.DB) (pb.ServicesServer, error) {
	return &servicesServer{&pb.ServicesDefaultServer{DB: database}}, nil
}

type servicesServer struct {
	*pb.ServicesDefaultServer
}
