package domain

import (
	"time"
)

// School represents a school entity
type School struct {
	ID          int
	Name        string
	Description string
	Address     string
	PhoneNumber string
	OwnerEmail  string
	Approved    bool
	Suspended   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewSchool creates a new School instance
func NewSchool(name, description, address, phone_number, ownerEmail string) *School {
	return &School{
		Name:        name,
		Description: description,
		Address:     address,
		PhoneNumber: phone_number,
		OwnerEmail:  ownerEmail,
		Approved:    false,
		Suspended:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
