package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		fmt.Println(token)
		// newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)
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

	dbSQL, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, err
	}
	defer dbSQL.Close()
	if err := migrate.MigrateDB(*dbSQL); err != nil {
		return nil, err
	}

	db, err := gorm.Open("postgres", dbSQL)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	// register all of our services into the grpcServer
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

	cs, err := svc.NewContactsServer(db)
	if err != nil {
		return nil, err
	}

	pb.RegisterContactsServer(grpcServer, cs)

	return grpcServer, nil
}
