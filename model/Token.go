package model

type AuthenticationToken struct {
	Id          int    `json:"id"`
	AccessToken string `json:"accessToken"`
}
