package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/user"
)

// UserRepository represents the repository interface for the User entity
type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

// PostgresUserRepository represents the PostgreSQL repository implementation for the User entity
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository creates a new instance of PostgresUserRepository
func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// Create creates a new user record in the database
func (r *PostgresUserRepository) Create(user *domain.User) error {
	// Implement the database insertion logic here
	return nil
}

// GetByID retrieves a user by ID from the database
func (r *PostgresUserRepository) GetByID(id int) (*domain.User, error) {
	// Implement the database retrieval logic here
	return nil, nil
}

// GetByEmail retrieves a user by email from the database
func (r *PostgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	// Implement the database retrieval logic here
	return nil, nil
}
