package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewGroupsServer(database *gorm.DB) (pb.GroupsServer, error) {
	return &groupsServer{&pb.GroupsDefaultServer{DB: database}}, nil
}

type groupsServer struct {
	*pb.GroupsDefaultServer
}
