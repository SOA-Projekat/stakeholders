package repo

import (
	"errors"

	"gorm.io/gorm"
	"stakeholders.xws.com/model"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (r *UserRepository) Get(userID int) (*model.User, error) {
	var user model.User
	result := r.DatabaseConnection.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *UserRepository) Update(user *model.User) (*model.User, error) {
	result := repo.DatabaseConnection.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		// Handle the case where no rows were updated, which might mean the user wasn't found
		return nil, errors.New("user does not exist")
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	result := r.DatabaseConnection.Find(&users)
	return users, result.Error
}
