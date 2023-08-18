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

func (s *VerificationService) Create(email string, token string, expiresAt int64) error {
	err := s.verificationRepo.Create(email, token, expiresAt)
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

func (s *VerificationService) Delete(token string) error {
	err := s.verificationRepo.Delete(token)
	if err != nil {
		return err
	}

	return nil
}
