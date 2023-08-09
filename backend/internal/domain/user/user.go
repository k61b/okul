package domain

import (
	"time"
)

// User represents a user entity
type User struct {
	ID           int
	Email        string
	PasswordHash string
	Name         string
	Surname      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new User instance
func NewUser(email, passwordHash, name, surname string) *User {
	return &User{
		Email:        email,
		PasswordHash: passwordHash,
		Name:         name,
		Surname:      surname,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

// GetFullName returns the user's full name
func (u *User) GetFullName() string {
	return u.Name + " " + u.Surname
}
