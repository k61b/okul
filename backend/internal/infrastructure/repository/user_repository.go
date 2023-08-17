package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/user"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id int) error
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	query := `
		INSERT INTO users (email, is_email_verified, password_hash, name, surname, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		user.Email,
		user.IsEmailVerified,
		user.Password,
		user.Name,
		user.Surname,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetByID(id int) (*domain.User, error) {
	query := `
		SELECT id, email, is_email_verified, password_hash, name, surname, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &domain.User{}

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.IsEmailVerified,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	query := `
        SELECT id, email, is_email_verified, password_hash, name, surname, created_at, updated_at
        FROM users
        WHERE email = $1
    `
	user := &domain.User{}

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.IsEmailVerified,
		&user.Password,
		&user.Name,
		&user.Surname,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) Update(user *domain.User) error {
	query := `
		UPDATE users
		SET name = $1, surname = $2, updated_at = $3
		WHERE id = $4
	`

	_, err := r.db.Exec(
		query,
		user.Name,
		user.Surname,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
