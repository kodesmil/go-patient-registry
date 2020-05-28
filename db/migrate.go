package db

import (
	"github.com/kodesmil/ks-backend/pkg/pb"
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func MigrateDB() error {
	db, err := gorm.Open("postgres", viper.GetString("database.dsn"))
	if err != nil {
		return err
	}
	defer db.Close()
	return db.AutoMigrate(
		&pb.ProfileORM{},
		&pb.GroupORM{},
		&pb.JournalSubjectORM{},
		&pb.JournalEntryORM{},
		&pb.FeedArticleORM{},
		&pb.FeedArticleDetailORM{},
		&pb.FeedAuthorORM{},
		&pb.FeedTagORM{},
		&pb.NotificationDeviceORM{},
		&pb.NotificationSettingORM{},
		&pb.ChatMessageORM{},
		&pb.ChatRoomORM{},
		&pb.ChatRoomParticipantORM{},
		&pb.LogActivityORM{},
	).Error
}
