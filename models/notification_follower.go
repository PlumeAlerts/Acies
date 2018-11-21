package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type NotificationFollower struct {
	NotificationsID int64  `gorm:"column:notifications_id;primary_key" json:"notifications_id"`
	UserID          string `gorm:"column:user_id" json:"user_id"`
	Username        string `gorm:"column:username" json:"username"`
	DisplayName     string `gorm:"column:display_name" json:"display_name"`
}

// TableName sets the insert table name for this struct type
func (n *NotificationFollower) TableName() string {
	return "notification_followers"
}
