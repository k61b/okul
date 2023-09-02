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

func (s *SchoolService) CreateSchool(name, description, address, phone_number, owner_email string) error {
	school := domain.NewSchool(name, description, address, phone_number, owner_email)

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

func (s *SchoolService) UpdateSchool(school *domain.School) error {
	err := s.schoolRepo.UpdateSchool(school)
	if err != nil {
		return err
	}

	return nil
}

func (s *SchoolService) SuspendSchool(id int) error {
	err := s.schoolRepo.SuspendSchool(id)
	if err != nil {
		return err
	}

	return nil
}
