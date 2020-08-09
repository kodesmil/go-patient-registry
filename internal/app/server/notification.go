package server

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"time"
)

type NotificationResult struct {
	DeviceToken string
}

func NewNotificationServer(dbConnectionString string) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("Error initializing app: %v\n", err)
	}

	c := cron.New()
	id, err := c.AddFunc("*/1 * * * *", func() {

		db, err := gorm.Open("postgres", dbConnectionString)
		defer db.Close()
		hours, minutes, _ := time.Now().Clock()
		pattern := fmt.Sprintf(
			"(\\*|%d) (\\*|%d) \\* \\* \\* \\*",
			hours,
			minutes,
		)

		ctx := context.Background()

		var notificationDevices []pb.NotificationDeviceORM

		err = db.
			Joins("left join notification_settings ns on ns.profile_id = notification_devices.profile_id").
			Where("ns.enable_notifications = ?", true).
			Where("ns.enable_journal_reminder = ?", true).
			Where("ns.cron_journal_reminder ~ ?", pattern).
			Find(&notificationDevices).
			Error

		if err != nil {
			log.Fatal(err)
		}

		if len(notificationDevices) == 0 {
			log.Infof("No notifications are about to be sent")
			return
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			log.Fatalf("error getting messaging client: %v", err)
		}

		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: "Hi",
				Body:  "Your favourite app is loving you  💖",
			},
			Tokens: pb.MapToDeviceTokens(notificationDevices),
		}

		// Send a message to the device corresponding to the provided
		// registration token.
		response, err := client.SendMulticast(ctx, message)
		if err != nil {
			log.Fatal(err)
		}

		// Response is a message ID string.
		log.Println("Successfully sent message:", response)
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Started job with id: ", id)
	c.Start()
}
