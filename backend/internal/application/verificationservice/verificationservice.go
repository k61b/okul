package verificationservice

import (
	"time"

	domain "github.com/k61b/okul/internal/domain/verification"
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type VerificationService struct {
	verificationRepo repository.VerificationRepository
}

func NewVerificationService(verificationRepo repository.VerificationRepository) *VerificationService {
	return &VerificationService{verificationRepo: verificationRepo}
}

func (s *VerificationService) CreateVerification(email, token string, expiresAt time.Time) error {
	verification := domain.NewVerification(email, token, expiresAt)

	err := s.verificationRepo.Create(verification)
	if err != nil {
		return err
	}

	return nil
}

func (s *VerificationService) GetEmailFromToken(token string) (string, error) {
	email, err := s.verificationRepo.GetEmailFromToken(token)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (s *VerificationService) DeleteByEmail(email string) error {
	err := s.verificationRepo.DeleteByEmail(email)
	if err != nil {
		return err
	}

	return nil
}
