package service

import (
	"errors"

	"stakeholders.xws.com/model"
	"stakeholders.xws.com/repo"
)

type AppRatingService struct {
	RatingRepo *repo.AppRatingRepository
}

func (service *AppRatingService) GetAll() ([]model.AppRating, error) {
	return service.RatingRepo.GetAll()
}

func (service *AppRatingService) Create(rating *model.AppRating) (*model.AppRating, error) {
	return service.RatingRepo.Create(rating)
}

func (service *AppRatingService) HasUserRated(userId int) (bool, error) {
	// Check if any AppRating exists with the given userId
	ratings, err := service.RatingRepo.GetAll()
	if err != nil {
		return false, errors.New("error getting ratings")
	}
	for _, rating := range ratings {
		if rating.UserId == userId {
			return true, nil
		}
	}
	return false, nil
}
