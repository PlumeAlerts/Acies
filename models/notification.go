package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type Notification struct {
	ID                 int64     `gorm:"column:id;primary_key" json:"id"`
	UserID             string    `gorm:"column:user_id" json:"user_id"`
	NotificationTypeID int       `gorm:"column:notification_type_id" json:"notification_type_id"`
	Notified           bool      `gorm:"column:notified;DEFAULT:false" json:"notified"`
	CreatedAt          time.Time `gorm:"column:created_at;DEFAULT:now()" json:"created_at"`
}

// TableName sets the insert table name for this struct type
func (n *Notification) TableName() string {
	return "notifications"
}
