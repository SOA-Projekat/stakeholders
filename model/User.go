package model

import "errors"

type UserRole int

const (
	Administrator UserRole = iota
	Author
	Tourist
)

var UserRoleStrings = map[UserRole]string{
	Administrator: "administrator",
	Author:        "author",
	Tourist:       "tourist",
}

type User struct {
	ID                int      `json:"id" gorm:"primaryKey"`
	Username          string   `json:"username" gorm:"unique"`
	Password          string   `json:"password"`
	Role              UserRole `json:"role"`
	IsActive          bool     `json:"isActive"`
	VerificationToken string   `json:"verificatonToken"`
}

func NewUser(username, password string, role UserRole, isActive bool, verificationToken string) (*User, error) {

	if username == "" {
		return nil, errors.New("invalid username")
	}
	if password == "" {
		return nil, errors.New("invalid password")
	}

	return &User{
		Username:          username,
		Password:          password,
		Role:              role,
		IsActive:          isActive,
		VerificationToken: verificationToken,
	}, nil
}
