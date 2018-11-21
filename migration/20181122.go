package migration

import (
	"github.com/PlumeAlerts/Acies/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate20181122() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "2018112200240",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(
				&models.User{},
				&models.Scope{},
				&models.UserToken{},
				&models.UserTokenScope{},
				&models.NotificationType{},
				&models.Notification{},
				&models.NotificationSub{},
				&models.NotificationGiftSub{},
				&models.NotificationBit{},
				&models.NotificationFollower{},
			).Error

			if err = tx.Model(models.UserToken{}).AddForeignKey("user_id", "users (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.UserTokenScope{}).AddForeignKey("user_token_id", "user_tokens (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.UserTokenScope{}).AddForeignKey("scope_id", "scopes (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.Notification{}).AddForeignKey("user_id", "users (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.Notification{}).AddForeignKey("notification_type_id", "notification_type (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.NotificationSub{}).AddForeignKey("notifications_id", "notifications (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.NotificationGiftSub{}).AddForeignKey("notifications_id", "notifications (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.NotificationBit{}).AddForeignKey("notifications_id", "notifications (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			if err = tx.Model(models.NotificationFollower{}).AddForeignKey("notifications_id", "notifications (id)", "RESTRICT", "RESTRICT").Error; err != nil {
				return err
			}
			return err
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable(
				&models.User{},
				&models.Scope{},
				&models.UserToken{},
				&models.UserTokenScope{},
				&models.NotificationType{},
				&models.Notification{},
				&models.NotificationSub{},
				&models.NotificationGiftSub{},
				&models.NotificationBit{},
				&models.NotificationFollower{},
			).Error
		},
	}
}
