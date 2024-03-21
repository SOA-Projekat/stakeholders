package model

import "time"

type AppRating struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserId      int       `json:"userId" gorm:"unique"`
	Rating      int       `json:"rating"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"dateCreated"`
}

func NewAppRating(userID, rating int, description string, dateCreated time.Time) *AppRating {
	return &AppRating{
		UserId:      userID,
		Rating:      rating,
		Description: description,
		DateCreated: dateCreated,
	}
}
