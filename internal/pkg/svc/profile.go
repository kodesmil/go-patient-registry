package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
	"golang.org/x/net/context"
)

func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
}

func (s *profilesServer) Create(ctx context.Context, in *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	return s.ProfilesDefaultServer.Create(ctx, in)
}

type profilesServer struct {
	*pb.ProfilesDefaultServer
}
