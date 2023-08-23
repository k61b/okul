package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/verification"
)

type VerificationRepository interface {
	Create(verification *domain.Verification) error
	GetVerificationInfoFromToken(token string) (int, string, string, error)
	DeleteByID(id int) error
}

type PostgresVerificationRepository struct {
	db *sql.DB
}

func NewPostgresVerificationRepository(db *sql.DB) *PostgresVerificationRepository {
	return &PostgresVerificationRepository{db: db}
}

func (r *PostgresVerificationRepository) Create(verification *domain.Verification) error {
	query := `
		INSERT INTO verification_tokens (type, email, token, expires_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		verification.VerificationType,
		verification.Email,
		verification.Token,
		verification.ExpiresAt,
	).Scan(&verification.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresVerificationRepository) GetVerificationInfoFromToken(token string) (int, string, string, error) {
	query := `
		SELECT id, type, email
		FROM verification_tokens
		WHERE token = $1
	`

	var id int
	var verificationType string
	var email string

	err := r.db.QueryRow(query, token).Scan(&id, &verificationType, &email)
	if err != nil {
		return 0, "", "", err
	}

	return id, verificationType, email, nil
}

func (r *PostgresVerificationRepository) DeleteByID(id int) error {
	query := `
		DELETE FROM verification_tokens
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
