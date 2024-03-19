package model

type Person struct {
	ID           int     `json:"id" gorm:"primaryKey"`
	UserId       int     `json:"userId" gorm:"unique"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Email        string  `json:"email"`
	ProfileImage string  `json:"profileImage"`
	Bio          string  `json:"bio"`
	Quote        string  `json:"quote"`
	Balance      float64 `json:"balance"`
}

func NewPerson(userID int, name, surname, email, profileImage, bio, quote string, balance float64) *Person {
	return &Person{
		UserId:       userID,
		Name:         name,
		Surname:      surname,
		Email:        email,
		ProfileImage: profileImage,
		Bio:          bio,
		Quote:        quote,
		Balance:      balance,
	}
}
