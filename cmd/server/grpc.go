package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/qor/admin"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kodesmil/atlas-app-toolkit/gateway"
	"github.com/kodesmil/atlas-app-toolkit/requestid"
	migrate "github.com/kodesmil/go-patient-registry/db"
	"github.com/kodesmil/go-patient-registry/pkg/pb"
	"github.com/kodesmil/go-patient-registry/pkg/svc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(logger *logrus.Logger, dbConnectionString string) (*grpc.Server, error) {

	firebaseAuth := func(ctx context.Context) (context.Context, error) {
		rawToken, err := grpc_auth.AuthFromMD(ctx, "Bearer")
		if err != nil {
			return nil, err
		}
		parser := jwt.Parser{}
		token, _, err := parser.ParseUnverified(rawToken, jwt.MapClaims{})
		if err != nil {
			return ctx, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ctx, errors.New("Error retrieving claim")
		}
		userID, ok := claims["user_id"]
		if !ok {
			return ctx, errors.New("Error retrieving claim")
		}
		var key interface{} = "AccountID"
		ctx = context.WithValue(ctx, key, userID)
		return ctx, nil
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    time.Duration(viper.GetInt("config.keepalive.time")) * time.Second,
				Timeout: time.Duration(viper.GetInt("config.keepalive.timeout")) * time.Second,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// authenticate
				grpc_auth.UnaryServerInterceptor(firebaseAuth),

				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),

				// Metrics middleware
				grpc_prometheus.UnaryServerInterceptor,

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)

	if err := migrate.MigrateDB(); err != nil {
		return nil, err
	}
	db, err := gorm.Open("postgres", dbConnectionString)
	// defer db.Close()

	ps, err := svc.NewProfilesServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterProfilesServer(grpcServer, ps)

	gs, err := svc.NewGroupsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterGroupsServer(grpcServer, gs)

	jss, err := svc.NewJournalSubjectsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterJournalSubjectsServer(grpcServer, jss)

	jse, err := svc.NewJournalEntriesServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterJournalEntriesServer(grpcServer, jse)

	fa, err := svc.NewFeedArticlesServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterFeedArticlesServer(grpcServer, fa)

	go func() {
		// Initialize
		Admin := admin.New(&admin.AdminConfig{DB: db})

		// Allow to use Admin to manage User, Product
		Admin.AddResource(&pb.JournalEntryORM{})
		Admin.AddResource(&pb.JournalSubjectORM{})
		Admin.AddResource(&pb.FeedTagORM{})
		article := Admin.AddResource(&pb.FeedArticleORM{})
		Admin.AddResource(&pb.FeedAuthorORM{})
		Admin.AddResource(&pb.ProfileORM{})
		Admin.AddResource(&pb.GroupORM{})
		article.Meta(&admin.Meta{Name: "Content", Type: "text"})
		mux := http.NewServeMux()

		Admin.MountTo("/admin", mux)

		fmt.Println("Listening on: 9000")
		err = http.ListenAndServe(":9000", mux)
		if err != nil {
			return
		}
	}()

	return grpcServer, nil
}
