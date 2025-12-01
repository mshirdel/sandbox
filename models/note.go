package models

import (
	"time"

	"gorm.io/gorm"
)

// Note represents a note in the system
type Note struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null" validate:"required"`
	Content   string         `json:"content" gorm:"type:text" validate:"required"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName returns the table name for the Note model
func (Note) TableName() string {
	return "notes"
}
