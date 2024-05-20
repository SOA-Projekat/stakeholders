package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"stakeholders.xws.com/model"
	"stakeholders.xws.com/repo"
)

type AccountService struct {
	PersonRepo *repo.PersonRepository
	UserRepo   *repo.UserRepository
}

func (service *AccountService) GenerateToken(userId int, username string, role model.UserRole) (*model.AuthenticationToken, error) {
	key := "explorer_secret_key"
	issuer := "explorer"
	audience := "explorer-front.com"

	userClaims := jwt.MapClaims{
		"jti":      uuid.New().String(),
		"id":       strconv.FormatInt(int64(userId), 10),
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Minute * 60 * 24 * 100).Unix(),
		"iss":      issuer,
		"aud":      audience,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	token, err := accessToken.SignedString([]byte(key))
	if err != nil {
		return nil, err
	}

	returnTokenValue := model.AuthenticationToken{
		Id:          userId,
		AccessToken: token,
	}

	return &returnTokenValue, nil

}

func (service *AccountService) CreateUser(user *model.User) (*model.AuthenticationToken, error) {

	if user.Username == "" || user.Password == "" {
		return nil, errors.New("username and password are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hashing password")
	}

	user, err = model.NewUser(user.Username, string(hashedPassword), user.Role, true, "")
	if err != nil {
		return nil, errors.New("error creating user model")
	}

	savedUser, err := service.UserRepo.Save(user)
	if err != nil {
		return nil, errors.New("error saving user to the database")
	}

	token, err := service.GenerateToken(savedUser.ID, savedUser.Username, savedUser.Role)
	if err != nil {
		return nil, errors.New("error generating token")
	}

	return token, nil
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
