package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type NotificationBit struct {
	NotificationsID  int64  `gorm:"column:notifications_id;primary_key" json:"notifications_id"`
	UserID           string `gorm:"column:user_id" json:"user_id"`
	Username         string `gorm:"column:username" json:"username"`
	DisplayName      string `gorm:"column:display_name" json:"display_name"`
	BitUsed          int    `gorm:"column:bit_used" json:"bit_used"`
	TotalBitsUsed    int    `gorm:"column:total_bits_used" json:"total_bits_used"`
	Message          string `gorm:"column:message" json:"message"`
	MessageID        string `gorm:"column:message_id" json:"message_id"`
	BadgeEntitlement int    `gorm:"column:badge_entitlement" json:"badge_entitlement"`
}

// TableName sets the insert table name for this struct type
func (n *NotificationBit) TableName() string {
	return "notification_bits"
}
