package postgres

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/user"
	"github.com/k61b/okul/internal/infrastructure/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	// Implement the database insertion logic here
	return nil
}

func (r *PostgresUserRepository) GetByID(id int) (*domain.User, error) {
	// Implement the database retrieval logic here
	return nil, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	// Implement the database retrieval logic here
	return nil, nil
}
