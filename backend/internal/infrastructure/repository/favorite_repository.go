package repository

import "database/sql"

type FavoriteRepository interface {
	Create(userID, schoolID int) error
	Get(userID int) ([]int, error)
	Delete(userID, schoolID int) error
}

type PostgresFavoriteRepository struct {
	db *sql.DB
}

func NewPostgresFavoriteRepository(db *sql.DB) *PostgresFavoriteRepository {
	return &PostgresFavoriteRepository{db: db}
}

func (r *PostgresFavoriteRepository) Create(userID, schoolID int) error {
	_, err := r.db.Exec("INSERT INTO favorites (user_id, school_id) VALUES ($1, $2)", userID, schoolID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresFavoriteRepository) Get(userID int) ([]int, error) {
	rows, err := r.db.Query("SELECT school_id FROM favorites WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	var schoolIDs []int
	for rows.Next() {
		var schoolID int
		err := rows.Scan(&schoolID)
		if err != nil {
			return nil, err
		}

		schoolIDs = append(schoolIDs, schoolID)
	}

	return schoolIDs, nil
}

func (r *PostgresFavoriteRepository) Delete(userID, schoolID int) error {
	_, err := r.db.Exec("DELETE FROM favorites WHERE user_id = $1 AND school_id = $2", userID, schoolID)
	if err != nil {
		return err
	}

	return nil
}
