package userservice

import (
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Implement methods for user-related actions here
// For example: CreateUser, GetUserByID, UpdateUser, DeleteUser, etc.
