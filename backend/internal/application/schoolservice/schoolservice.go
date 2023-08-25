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

func (s *SchoolService) GetAllSchools() ([]*domain.School, error) {
	schools, err := s.schoolRepo.GetAllSchools()
	if err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *SchoolService) GetSchoolByID(id int) (*domain.School, error) {
	school, err := s.schoolRepo.GetSchoolByID(id)
	if err != nil {
		return nil, err
	}

	return school, nil
}
