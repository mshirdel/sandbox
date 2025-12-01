package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `json:"id" db:"id"`
	Email     string    `json:"email" db:"email" validate:"required,email"`
	Name      string    `json:"name" db:"name" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Message represents a message from a user
type Message struct {
	ID        uint      `json:"id" db:"id"`
	UserID    uint      `json:"user_id" db:"user_id" validate:"required"`
	Title     string    `json:"title" db:"title" validate:"required"`
	Content   string    `json:"content" db:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
