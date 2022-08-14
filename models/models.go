package models

import (
	"database/sql"
	"time"
)

// Base model struct
type Base struct {
	ID        uint         `json:"id,omitempty" gorm:"primary_key,autoIncrement,not null"`
	CreatedAt time.Time    `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"not null"`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"index,not null"`
}
