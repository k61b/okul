package schoolservice

import (
	domain "github.com/k61b/okul/internal/domain/school"
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type SchoolService struct {
	schoolRepo repository.SchoolRepository
}

func NewSchoolService(schoolRepo repository.SchoolRepository) *SchoolService {
	return &SchoolService{schoolRepo: schoolRepo}
}

func (s *SchoolService) CreateSchool(name, description, address, phone_number string, owner_id int) error {
	school := domain.NewSchool(name, description, address, phone_number, owner_id)

	err := s.schoolRepo.Create(school)
	if err != nil {
		return err
	}

	return nil
}
