package domain

import (
	"time"
)

// School represents a school entity
type School struct {
	ID        int
	Name      string
	Address   string
	OwnerID   int
	Approved  bool
	Suspended bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewSchool creates a new School instance
func NewSchool(name, address string, ownerID int) *School {
	return &School{
		Name:      name,
		Address:   address,
		OwnerID:   ownerID,
		Approved:  false,
		Suspended: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
