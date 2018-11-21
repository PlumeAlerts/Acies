package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type NotificationType struct {
	ID           int    `gorm:"column:id;primary_key" json:"id"`
	Notification string `gorm:"column:notification" json:"notification"`
}

// TableName sets the insert table name for this struct type
func (n *NotificationType) TableName() string {
	return "notification_type"
}
