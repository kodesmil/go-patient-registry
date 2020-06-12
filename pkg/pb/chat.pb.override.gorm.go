package pb

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	gorm1 "github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func MapToDeviceTokens(vs []NotificationDeviceORM) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = v.DeviceToken
	}
	return vsm
}

func (m *ChatMessageORM) AfterCreate_(ctx context.Context, db *gorm1.DB) error {
	go func() {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			logrus.Fatalf("error initializing app: %v\n", err)
		}
		var notificationDevices []NotificationDeviceORM
		err = db.
			Select("notification_devices.device_token").
			Joins("left join notification_settings ns on ns.account_id = notification_devices.account_id").
			Joins("left join chat_room_participants crp on crp.profile_id = ns.account_id").
			Where("ns.enable_notifications = ?", true).
			Where("crp.chat_room_id = (select chat_room_id from chat_room_participants where id = ?)", m.AuthorId).
			Find(&notificationDevices).
			Error

		if err != nil {
			logrus.Fatal(err)
		}

		if len(notificationDevices) == 0 {
			logrus.Info("No notifications to be sent")
			return
		}

		client, err := app.Messaging(ctx)
		if err != nil {
			logrus.Fatalf("error getting Messaging client: %v\n", err)
		}
		message := &messaging.MulticastMessage{
			Notification: &messaging.Notification{
				Title: fmt.Sprintf("New message"),
				Body:  m.Text,
			},
			Tokens: MapToDeviceTokens(notificationDevices),
		}
		// Send a message to the device corresponding to the provided
		// registration token.
		response, err := client.SendMulticast(ctx, message)
		if err != nil {
			logrus.Fatalln(err)
		}
		// Response is a message ID string.
		logrus.Println("Successfully sent message:", response)
	}()
	return nil
}
