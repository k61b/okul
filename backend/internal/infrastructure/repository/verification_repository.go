package repository

import (
	"database/sql"
)

type VerificationRepository interface {
	Create(email string, token string, expiresAt int64) error
	GetEmailFromToken(token string) (string, error)
	Delete(token string) error
}

type PostgresVerificationRepository struct {
	db *sql.DB
}

func NewPostgresVerificationRepository(db *sql.DB) *PostgresVerificationRepository {
	return &PostgresVerificationRepository{db: db}
}

func (r *PostgresVerificationRepository) Create(email string, token string, expiresAt int64) error {
	query := `
		INSERT INTO verifications (email, token, expires_at)	
		VALUES ($1, $2, $3)
	`
	_, err := r.db.Exec(query, email, token, expiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresVerificationRepository) GetEmailFromToken(token string) (string, error) {
	query := `
		SELECT email
		FROM verifications
		WHERE token = $1
	`
	var email string
	err := r.db.QueryRow(query, token).Scan(&email)
	if err != nil {
		return "", err
	}
	return email, nil
}

func (r *PostgresVerificationRepository) Delete(token string) error {
	query := `
		DELETE FROM verifications
		WHERE token = $1
	`
	_, err := r.db.Exec(query, token)
	if err != nil {
		return err
	}
	return nil
}
