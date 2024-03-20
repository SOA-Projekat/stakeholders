package service

import (
	"errors"

	"stakeholders.xws.com/model"
	"stakeholders.xws.com/repo"
)

type AccountService struct {
	PersonRepo *repo.PersonRepository
	UserRepo   *repo.UserRepository
}

func (service *AccountService) GetAll() ([]model.Account, error) {

	var accounts []model.Account

	users, err := service.UserRepo.GetAll()
	if err != nil {
		return accounts, errors.New("error getting users")
	}

	for _, user := range users {
		person, _ := service.PersonRepo.Get(user.ID)

		account := model.Account{
			UserId:   user.ID,
			Username: user.Username,
			Password: user.Password,
			Email:    person.Email,
			Role:     model.UserRoleStrings[user.Role],
			IsActive: user.IsActive,
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (service *AccountService) BlockOrUnblock(account *model.Account) (*model.Account, error) {
	user, err := service.UserRepo.Get(account.UserId)
	if err != nil {
		return nil, errors.New("error getting user")
	}

	if account.IsActive {
		user.IsActive = false
	} else {
		user.IsActive = true
	}

	user, err = service.UserRepo.Update(user)
	if err != nil {
		return nil, errors.New("error updating user")
	}

	account.IsActive = user.IsActive
	return account, nil
}
