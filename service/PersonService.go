package service

import (
	"fmt"

	"stakeholders.xws.com/model"
	"stakeholders.xws.com/repo"
)

type PersonService struct {
	PersonRepo *repo.PersonRepository
}

func (service *PersonService) FindPerson(userID int) (*model.Person, error) {
	person, err := service.PersonRepo.Get(userID)
	if err != nil {
		return nil, fmt.Errorf("person with user ID %d not found", userID)
	}
	return person, nil
}

func (service *PersonService) UpdatePerson(updatedPerson *model.Person) (*model.Person, error) {
	person, err := service.PersonRepo.Update(updatedPerson)
	if err != nil {
		return nil, err
	}
	return person, nil
}
