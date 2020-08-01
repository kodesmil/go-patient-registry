package db

import (
	"github.com/kodesmil/ks-backend/internal/pkg/pb"
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
		&pb.HealthMenstruationDailyEntryORM{},
		&pb.HealthMenstruationPersonalInfoORM{},
		&pb.FeedArticleORM{},
		&pb.FeedArticleDetailORM{},
		&pb.FeedAuthorORM{},
		&pb.FeedTagORM{},
		&pb.NotificationDeviceORM{},
		&pb.NotificationSettingORM{},
		&pb.ChatMessageORM{},
		&pb.ChatRoomORM{},
		&pb.ChatRoomParticipantORM{},
		&pb.ServiceORM{},
		&pb.ServiceProviderORM{},
		&pb.ServiceProviderSessionEvaluationORM{},
		&pb.ServiceTagORM{},
		&pb.ServiceSessionORM{},
		&pb.ServiceSessionNoteORM{},
		&pb.ServiceDetailsORM{},
		&pb.ServiceDetailsContactORM{},
		&pb.ServiceDetailsCompanyORM{},
		&pb.ServiceOfferORM{},
		&pb.ServiceApplicationORM{},
		&pb.ServiceApplicationFileORM{},
	).Error
}
