package svc

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
)

// NewProfilesServer returns an instance of the default profiles server interface
func NewProfilesServer(database *gorm.DB) (pb.ProfilesServer, error) {
	return &profilesServer{&pb.ProfilesDefaultServer{DB: database}}, nil
}

type profilesServer struct {
	*pb.ProfilesDefaultServer
}

// NewGroupsServer returns an instance of the default groups server interface
func NewGroupsServer(database *gorm.DB) (pb.GroupsServer, error) {
	return &groupsServer{&pb.GroupsDefaultServer{DB: database}}, nil
}

type groupsServer struct {
	*pb.GroupsDefaultServer
}

func (s *contactsServer) Create(ctx context.Context, in *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	println("Hello")
	fmt.Println(in)
	resp, _ := s.ContactsDefaultServer.Create(ctx, in)
	return resp, nil
}

func NewContactsServer(database *gorm.DB) (pb.ContactsServer, error) {
	return &contactsServer{&pb.ContactsDefaultServer{DB: database}}, nil
}

type contactsServer struct {
	*pb.ContactsDefaultServer
}
