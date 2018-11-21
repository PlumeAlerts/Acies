package models

import (
	"database/sql"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
)

type Scope struct {
	ID    int    `gorm:"column:id;primary_key" json:"id"`
	Scope string `gorm:"column:scope" json:"scope"`
}

// TableName sets the insert table name for this struct type
func (s *Scope) TableName() string {
	return "scopes"
}
