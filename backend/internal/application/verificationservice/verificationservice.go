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

func (s *VerificationService) CreateVerification(verificationType, email, token string, expiresAt time.Time) error {
	verification := domain.NewVerification(verificationType, email, token, expiresAt)

	err := s.verificationRepo.Create(verification)
	if err != nil {
		return err
	}

	return nil
}

func (s *VerificationService) GetVerificationInfoFromToken(token string) (int, string, string, error) {
	id, verificationType, email, err := s.verificationRepo.GetVerificationInfoFromToken(token)
	if err != nil {
		return 0, "", "", err
	}

	return id, verificationType, email, nil
}

func (s *VerificationService) DeleteVerification(id int) error {
	err := s.verificationRepo.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
