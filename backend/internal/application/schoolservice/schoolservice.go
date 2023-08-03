package schoolservice

import (
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type SchoolService struct {
	schoolRepo repository.SchoolRepository
}

func NewSchoolService(schoolRepo repository.SchoolRepository) *SchoolService {
	return &SchoolService{schoolRepo: schoolRepo}
}

// Implement methods for school-related actions here
// For example: CreateSchool, GetSchoolByID, UpdateSchool, etc.
