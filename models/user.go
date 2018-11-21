package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type User struct {
	ID          string `gorm:"column:id;primary_key" json:"id"`
	DisplayName string `gorm:"column:display_name" json:"display_name"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
