package main

import (
	"context"
	"errors"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/qor/admin"
	"github.com/robfig/cron/v3"
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

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		logrus.Fatalf("error initializing app: %v\n", err)
	}

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
		grpc.ChainStreamInterceptor(
			grpc_middleware.ChainStreamServer(
				// authenticate
				grpc_auth.StreamServerInterceptor(firebaseAuth),
				// logging middleware
				grpc_logrus.StreamServerInterceptor(logrus.NewEntry(logger)),
				// Request-Id interceptor
				requestid.StreamServerInterceptor(),
				// Metrics middleware
				grpc_prometheus.StreamServerInterceptor,
				// validation middleware
				grpc_validator.StreamServerInterceptor(),
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

	fad, err := svc.NewFeedArticleDetailsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterFeedArticleDetailsServer(grpcServer, fad)

	nd, err := svc.NewNotificationDevicesServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterNotificationDevicesServer(grpcServer, nd)

	ns, err := svc.NewNotificationSettingsServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterNotificationSettingsServer(grpcServer, ns)

	cm := svc.NewChatServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterChatServer(grpcServer, cm)

	go func() {
		// Initialize
		Admin := admin.New(&admin.AdminConfig{DB: db})

		// Allow to use Admin to manage User, Product
		Admin.AddResource(&pb.JournalEntryORM{})
		Admin.AddResource(&pb.JournalSubjectORM{})
		Admin.AddResource(&pb.FeedTagORM{})
		article := Admin.AddResource(&pb.FeedArticleORM{})
		article.Meta(&admin.Meta{Name: "Content", Type: "text"})
		articleDetail := Admin.AddResource(&pb.FeedArticleDetailORM{})
		articleDetail.Meta(&admin.Meta{Name: "Content", Type: "text"})
		Admin.AddResource(&pb.FeedAuthorORM{})
		Admin.AddResource(&pb.ProfileORM{})
		Admin.AddResource(&pb.GroupORM{})
		Admin.AddResource(&pb.NotificationSettingORM{})
		Admin.AddResource(&pb.NotificationDeviceORM{})
		Admin.AddResource(&pb.ChatMessageORM{})
		Admin.AddResource(&pb.LogActivityORM{})
		mux := http.NewServeMux()

		Admin.MountTo("/admin", mux)

		logrus.Println("Listening on: 9000")
		err = http.ListenAndServe(":9000", mux)
		if err != nil {
			return
		}
	}()

	type NotificationResult struct {
		DeviceToken string
	}

	go func() {

		c := cron.New()
		id, err := c.AddFunc("*/1 * * * *", func() {

			hours, minutes, _ := time.Now().Clock()
			pattern := fmt.Sprintf(
				"(\\*|%d) (\\*|%d) \\* \\* \\* \\*",
				hours,
				minutes,
			)

			ctx := context.Background()
			rows, err := db.Table("notification_settings ns").
				Select("nd.device_token").
				Joins("left join notification_devices nd on nd.account_id = ns.account_id").
				Where("ns.enable_notifications = ?", true).
				Where("ns.enable_journal_reminder = ?", true).
				Where("ns.cron_journal_reminder ~ ?", pattern).
				Rows()
			if err != nil {
				logrus.Println(err)
			}
			defer rows.Close()

			for rows.Next() {
				var result NotificationResult
				db.ScanRows(rows, &result)

				client, err := app.Messaging(ctx)
				if err != nil {
					logrus.Fatalf("error getting Messaging client: %v\n", err)
				}
				message := &messaging.Message{
					Notification: &messaging.Notification{
						Title: "Wake up! Wake up! Wake up!",
						Body:  "Have a great day  💖",
					},
					Token: result.DeviceToken,
				}

				// Send a message to the device corresponding to the provided
				// registration token.
				response, err := client.Send(ctx, message)
				if err != nil {
					logrus.Fatalln(err)
				}

				// Response is a message ID string.
				logrus.Println("Successfully sent message:", response)
			}

		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Started job with id: ", id)
		c.Start()
	}()

	return grpcServer, nil
}
