package svc

import (
	"fmt"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
	"golang.org/x/net/context"
)

func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
}

func (s *profilesServer) Create(ctx context.Context, in *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	in.Payload.Id = &resource.Identifier{ResourceId: fmt.Sprintf("%v", ctx.Value("AccountID"))}
	return s.ProfilesDefaultServer.Create(ctx, in)
}

type profilesServer struct {
	*pb.ProfilesDefaultServer
}
