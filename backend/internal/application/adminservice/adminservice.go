package adminservice

import (
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type AdminService struct {
	adminRepo repository.AdminRepository
}

func NewAdminService(adminRepo repository.AdminRepository) *AdminService {
	return &AdminService{adminRepo: adminRepo}
}

// Implement methods for admin-related actions here
// For example: CreateAdmin, GetAdminByID, DeleteAdmin, etc.
