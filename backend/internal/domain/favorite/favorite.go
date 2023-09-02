package domain

type Favorite struct {
	ID        int
	UserID    int
	SchoolID  int
	CreatedAt string
	UpdatedAt string
}

func NewFavorite(userID, schoolID int) *Favorite {
	return &Favorite{
		UserID:   userID,
		SchoolID: schoolID,
	}
}
