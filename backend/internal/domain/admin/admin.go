package domain

import (
	"time"
)

// Admin represents an admin entity
type Admin struct {
	ID           int
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewAdmin creates a new Admin instance
func NewAdmin(email, passwordHash string) *Admin {
	return &Admin{
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GetEmail returns the admin's email
func (a *Admin) GetEmail() string {
	return a.Email
}
