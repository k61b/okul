package favoriteservice

import "github.com/k61b/okul/internal/infrastructure/repository"

type FavoriteService struct {
	favoriteRepo repository.FavoriteRepository
}

func NewFavoriteService(favoriteRepo repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{favoriteRepo: favoriteRepo}
}

func (s *FavoriteService) CreateFavorite(userID, schoolID int) error {
	err := s.favoriteRepo.Create(userID, schoolID)
	if err != nil {
		return err
	}

	return nil
}

func (s *FavoriteService) GetFavoriteSchoolIDs(userID int) ([]int, error) {
	schoolIDs, err := s.favoriteRepo.Get(userID)
	if err != nil {
		return nil, err
	}

	return schoolIDs, nil
}

func (s *FavoriteService) DeleteFavorite(userID, schoolID int) error {
	err := s.favoriteRepo.Delete(userID, schoolID)
	if err != nil {
		return err
	}

	return nil
}
