package svc

import (
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/pkg/pb"
)

func NewFeedArticlesServer(database *gorm.DB) (pb.FeedArticlesServer, error) {
	return &feedArticlesServer{&pb.FeedArticlesDefaultServer{DB: database}}, nil
}

type feedArticlesServer struct {
	*pb.FeedArticlesDefaultServer
}

func NewFeedArticleDetailsServer(database *gorm.DB) (pb.FeedArticleDetailsServer, error) {
	return &feedArticleDetailsServer{&pb.FeedArticleDetailsDefaultServer{DB: database}}, nil
}

type feedArticleDetailsServer struct {
	*pb.FeedArticleDetailsDefaultServer
}
