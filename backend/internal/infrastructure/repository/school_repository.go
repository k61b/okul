package repository

import (
	"database/sql"

	domain "github.com/k61b/okul/internal/domain/school"
)

type SchoolRepository interface {
	Create(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
	GetSchoolByID(id int) (*domain.School, error)
	UpdateSchool(school *domain.School) error
	SuspendSchool(id int) error
}

type PostgresSchoolRepository struct {
	db *sql.DB
}

func NewPostgresSchoolRepository(db *sql.DB) *PostgresSchoolRepository {
	return &PostgresSchoolRepository{db: db}
}

func (r *PostgresSchoolRepository) Create(school *domain.School) error {
	query := `
		INSERT INTO schools (name, description, address, phone_number, owner_email, approved, suspended, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		school.Name,
		school.Description,
		school.Address,
		school.PhoneNumber,
		school.OwnerEmail,
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
		SELECT id, name, description, address, phone_number, owner_email, approved, suspended, created_at, updated_at
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
			&school.OwnerEmail,
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

func (r *PostgresSchoolRepository) GetSchoolByID(id int) (*domain.School, error) {
	query := `
		SELECT id, name, description, address, phone_number, owner_email, approved, suspended, created_at, updated_at
		FROM schools
		WHERE id = $1
	`

	school := &domain.School{}

	err := r.db.QueryRow(query, id).Scan(
		&school.ID,
		&school.Name,
		&school.Description,
		&school.Address,
		&school.PhoneNumber,
		&school.OwnerEmail,
		&school.Approved,
		&school.Suspended,
		&school.CreatedAt,
		&school.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return school, nil
}

func (r *PostgresSchoolRepository) UpdateSchool(school *domain.School) error {
	query := `
		UPDATE schools
		SET name = $1, description = $2, address = $3, phone_number = $4, updated_at = $5
		WHERE id = $6
	`

	_, err := r.db.Exec(
		query,
		school.Name,
		school.Description,
		school.Address,
		school.PhoneNumber,
		school.UpdatedAt,
		school.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresSchoolRepository) SuspendSchool(id int) error {
	query := `
		UPDATE schools
		SET suspended = true
		WHERE id = $1
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
