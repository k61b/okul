package postgres

import (
	"database/sql"

	repository "github.com/k61b/okul/internal/infrastructure/repository"
	_ "github.com/lib/pq"
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

func (p *PostgreSQLDB) Close() error {
	return p.db.Close()
}

func (p *PostgreSQLDB) DB() *sql.DB {
	return p.db
}

func (p *PostgreSQLDB) UserRepo() repository.UserRepository {
	return repository.NewPostgresUserRepository(p.db)
}

func (p *PostgreSQLDB) SchoolRepo() repository.SchoolRepository {
	return repository.NewPostgresSchoolRepository(p.db)
}

func (p *PostgreSQLDB) VerificationRepo() repository.VerificationRepository {
	return repository.NewPostgresVerificationRepository(p.db)
}
