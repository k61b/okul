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

// Approve approves the school
func (s *School) Approve() {
	s.Approved = true
	s.UpdatedAt = time.Now()
}

// Suspend suspends the school
func (s *School) Suspend() {
	s.Suspended = true
	s.UpdatedAt = time.Now()
}

// Resume resumes the school
func (s *School) Resume() {
	s.Suspended = false
	s.UpdatedAt = time.Now()
}

// IsApproved returns true if the school is approved
func (s *School) IsApproved() bool {
	return s.Approved
}

// IsSuspended returns true if the school is suspended
func (s *School) IsSuspended() bool {
	return s.Suspended
}

// UpdateName updates the school's name
func (s *School) UpdateName(name string) {
	s.Name = name
	s.UpdatedAt = time.Now()
}

// UpdateAddress updates the school's address
func (s *School) UpdateAddress(address string) {
	s.Address = address
	s.UpdatedAt = time.Now()
}
