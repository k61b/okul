package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/verification"
)

type VerificationRepository interface {
	Create(verification *domain.Verification) error
	GetEmailFromToken(token string) (string, error)
	DeleteByEmail(email string) error
}

type PostgresVerificationRepository struct {
	db *sql.DB
}

func NewPostgresVerificationRepository(db *sql.DB) *PostgresVerificationRepository {
	return &PostgresVerificationRepository{db: db}
}

func (r *PostgresVerificationRepository) Create(verification *domain.Verification) error {
	query := `
		INSERT INTO verification_tokens (email, token, expires_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		verification.Email,
		verification.Token,
		verification.ExpiresAt,
	).Scan(&verification.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresVerificationRepository) GetEmailFromToken(token string) (string, error) {
	query := `
		SELECT email
		FROM verification_tokens
		WHERE token = $1
	`

	var email string

	err := r.db.QueryRow(query, token).Scan(&email)
	if err != nil {
		return "", err
	}

	return email, nil
}

func (r *PostgresVerificationRepository) DeleteByEmail(email string) error {
	query := `
		DELETE FROM verification_tokens
		WHERE email = $1
	`

	_, err := r.db.Exec(query, email)
	if err != nil {
		return err
	}

	return nil
}
