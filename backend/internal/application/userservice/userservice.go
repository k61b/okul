package userservice

import (
	domain "github.com/k61b/okul/internal/domain/user"
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(email, password, name, surname string) (*domain.User, error) {
	user := domain.NewUser(email, password, name, surname)

	err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByID(id int) (*domain.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Update(user *domain.User) (*domain.User, error) {
	err := s.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Delete(id int) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
