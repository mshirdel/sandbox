package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Username  string         `gorm:"column:username;type:VARCHAR(128);not null;unique"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
