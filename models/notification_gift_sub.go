package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type NotificationGiftSub struct {
	NotificationsID   int64  `gorm:"column:notifications_id;primary_key" json:"notifications_id"`
	UserID            string `gorm:"column:user_id" json:"user_id"`
	Username          string `gorm:"column:username" json:"username"`
	DisplayName       string `gorm:"column:display_name" json:"display_name"`
	SubPlan           string `gorm:"column:sub_plan" json:"sub_plan"`
	Months            int    `gorm:"column:months" json:"months"`
	GifterID          string `gorm:"column:gifter_id" json:"gifter_id"`
	GifterUsername    string `gorm:"column:gifter_username" json:"gifter_username"`
	GifterDisplayName string `gorm:"column:gifter_display_name" json:"gifter_display_name"`
}

// TableName sets the insert table name for this struct type
func (n *NotificationGiftSub) TableName() string {
	return "notification_gift_subs"
}
