package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/school"
)

type SchoolRepository interface {
	Create(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
}

type PostgresSchoolRepository struct {
	db *sql.DB
}

func NewPostgresSchoolRepository(db *sql.DB) *PostgresSchoolRepository {
	return &PostgresSchoolRepository{db: db}
}

func (r *PostgresSchoolRepository) Create(school *domain.School) error {
	query := `
		INSERT INTO schools (name, description, address, phone_number, owner_id, approved, suspended, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		school.Name,
		school.Description,
		school.Address,
		school.PhoneNumber,
		school.OwnerID,
		school.Approved,
		school.Suspended,
		school.CreatedAt,
		school.UpdatedAt,
	).Scan(&school.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresSchoolRepository) GetAllSchools() ([]*domain.School, error) {
	query := `
		SELECT id, name, description, address, phone_number, owner_id, approved, suspended, created_at, updated_at
		FROM schools
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	schools := make([]*domain.School, 0)

	for rows.Next() {
		school := &domain.School{}
		err := rows.Scan(
			&school.ID,
			&school.Name,
			&school.Description,
			&school.Address,
			&school.PhoneNumber,
			&school.OwnerID,
			&school.Approved,
			&school.Suspended,
			&school.CreatedAt,
			&school.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}

	return schools, nil
}
