package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewPeriodServer(database *gorm.DB) (pb.PeriodServer, error) {
	return &periodServer{&pb.PeriodDefaultServer{DB: database}}, nil
}

type periodServer struct {
	*pb.PeriodDefaultServer
}
