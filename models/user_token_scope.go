package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type UserTokenScope struct {
	UserTokenID int64 `gorm:"column:user_token_id;primary_key" json:"user_token_id"`
	ScopeID     int   `gorm:"column:scope_id" json:"scope_id"`
}

// TableName sets the insert table name for this struct type
func (u *UserTokenScope) TableName() string {
	return "user_token_scope"
}
