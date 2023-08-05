package domain

import (
	"time"
)

// User represents a user entity
type User struct {
	ID        int
	Email     string
	Password  string // This should be a hashed password
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new User instance
func NewUser(email, password string, isAdmin bool) *User {
	return &User{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
