package service

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"stakeholders.xws.com/model"
)

type JWTService struct {
}

func (service *JWTService) GenerateToken(userId int, username string, role model.UserRole) (*model.AuthenticationToken, error) {
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
