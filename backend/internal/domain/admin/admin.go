package domain

import (
	"time"
)

// Admin represents a user entity
type Admin struct {
	ID        int
	Email     string
	Password  string // This should be a hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new Admin instance
func NewAdmin(email, password string, isAdmin bool) *Admin {
	return &Admin{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
