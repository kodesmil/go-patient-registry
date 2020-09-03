package svc

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
)

type servicesServer struct {
	*pb.ServicesDefaultServer
}

func NewServicesServer(database *gorm.DB) (pb.ServicesServer, error) {
	return &servicesServer{&pb.ServicesDefaultServer{DB: database}}, nil
}

func (s *servicesServer) ListServiceSession(ctx context.Context, in *pb.ListServiceSessionRequest) (*pb.ListServiceSessionResponse, error) {
	return s.CustomListServiceSession(ctx, in)
}

func (s *servicesServer) ListServiceOfferSession(ctx context.Context, in *pb.ListServiceOfferSessionRequest) (*pb.ListServiceOfferSessionResponse, error) {
	return s.CustomListServiceOfferSession(ctx, in)
}
