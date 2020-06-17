package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewHealthServer(database *gorm.DB) (pb.HealthServer, error) {
	return &healthServer{&pb.HealthDefaultServer{DB: database}}, nil
}

type healthServer struct {
	*pb.HealthDefaultServer
}
