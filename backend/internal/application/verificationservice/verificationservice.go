package verificationservice

import (
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type VerificationService struct {
	verificationRepo repository.VerificationRepository
}

func NewVerificationService(verificationRepo repository.VerificationRepository) *VerificationService {
	return &VerificationService{verificationRepo: verificationRepo}
}

func (s *VerificationService) Create(email string, token string, expiresAt string) error {
	err := s.verificationRepo.Create(email, token, expiresAt)
	if err != nil {
		return err
	}

	return nil
}
