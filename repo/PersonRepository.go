package repo

import (
	"gorm.io/gorm"
	"stakeholders.xws.com/model"
)

// PersonRepository handles database operations for Person entities.
type PersonRepository struct {
	DatabaseConnection *gorm.DB
}

// Get retrieves a person record by userID from the database.
func (r *PersonRepository) Get(userID int) (*model.Person, error) {
	var person model.Person
	result := r.DatabaseConnection.First(&person, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &person, nil
}

func (r *PersonRepository) Update(updatedPerson *model.Person) (*model.Person, error) {
	// Check if the person exists in the database
	var existingPerson model.Person
	result := r.DatabaseConnection.First(&existingPerson, "user_id = ?", updatedPerson.UserId)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update fields of the found person record
	existingPerson.Name = updatedPerson.Name
	existingPerson.Surname = updatedPerson.Surname
	existingPerson.Email = updatedPerson.Email
	existingPerson.ProfileImage = updatedPerson.ProfileImage
	existingPerson.Bio = updatedPerson.Bio
	existingPerson.Quote = updatedPerson.Quote
	existingPerson.Balance = updatedPerson.Balance
	// Save the updated person record
	result = r.DatabaseConnection.Save(&existingPerson)
	if result.Error != nil {
		return nil, result.Error
	}

	return &existingPerson, nil
}
