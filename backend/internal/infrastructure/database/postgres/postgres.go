package postgres

import (
	"database/sql"

	"github.com/k61b/okul/internal/infrastructure/repository"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type PostgreSQLDB struct {
	db *sql.DB
}

func NewPostgreSQLDB(connectionString string) (*PostgreSQLDB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgreSQLDB{db: db}, nil
}

func (p *PostgreSQLDB) UserRepo() repository.UserRepository {
	return NewPostgresUserRepository(p.db)
}

func (p *PostgreSQLDB) SchoolRepo() repository.SchoolRepository {
	return NewPostgresSchoolRepository(p.db)
}
