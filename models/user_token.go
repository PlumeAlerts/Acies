package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type UserToken struct {
	ID           int64 `gorm:"column:id;primary_key" json:"id"`
	User         User
	UserID       string    `gorm:"column:user_id" json:"user_id"`
	AccessToken  string    `gorm:"column:access_token" json:"access_token"`
	RefreshToken string    `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiresIn    time.Time `gorm:"column:expires_in" json:"expires_in"`
}

// TableName sets the insert table name for this struct type
func (u *UserToken) TableName() string {
	return "user_tokens"
}
