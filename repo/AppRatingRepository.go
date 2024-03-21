package repo

import (
	"gorm.io/gorm"
	"stakeholders.xws.com/model"
)

// AppRatingRepository handles database operations for AppRating entities.
type AppRatingRepository struct {
	DatabaseConnection *gorm.DB
}

// GetAll retrieves all AppRating records from the database.
func (r *AppRatingRepository) GetAll() ([]model.AppRating, error) {
	var ratings []model.AppRating
	result := r.DatabaseConnection.Find(&ratings)
	return ratings, result.Error
}

// Create creates a new AppRating record in the database.
func (r *AppRatingRepository) Create(rating *model.AppRating) (*model.AppRating, error) {
	dbCreationResult := r.DatabaseConnection.Create(rating)
	if dbCreationResult.Error != nil {
		return nil, dbCreationResult.Error
	}
	return rating, nil
}
