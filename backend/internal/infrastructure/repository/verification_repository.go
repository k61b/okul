package repository

import (
	"database/sql"
)

type VerificationRepository interface {
	Create(email string, token string, expiresAt string) error
	GetByEmail(email string) (string, error)
	Delete(email string) error
}

type PostgresVerificationRepository struct {
	db *sql.DB
}

func NewPostgresVerificationRepository(db *sql.DB) *PostgresVerificationRepository {
	return &PostgresVerificationRepository{db: db}
}

func (r *PostgresVerificationRepository) Create(email string, token string, expiresAt string) error {
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

func (r *PostgresVerificationRepository) GetByEmail(email string) (string, error) {
	query := `
		SELECT token
		FROM verifications
		WHERE email = $1
	`
	var token string
	err := r.db.QueryRow(query, email).Scan(&token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *PostgresVerificationRepository) Delete(email string) error {
	query := `
		DELETE FROM verifications
		WHERE email = $1
	`
	_, err := r.db.Exec(query, email)
	if err != nil {
		return err
	}
	return nil
}
