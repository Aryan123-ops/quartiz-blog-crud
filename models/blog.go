package models

import (
	"time"
)

// Creates the table blogs in the database.
type Blog struct {
	// gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"descriptions"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}

// gorm Model automatically creates fields like ID, Created_at, updated_at automatically
