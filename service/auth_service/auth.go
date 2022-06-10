package auth_service

import "github.com/jameschun7/go-gin-sample/models"

type Auth struct {
	UserID      string
	Username    string
	LoginType   string
	Country     string
	Email       string
	UserPicture string
}

// func (a *Auth) Check() (bool, error) {
// 	return models.CheckAuth(a.Username, a.Password)
// }

func (a *Auth) ExistByUserID() (*models.Account, error) {
	return models.ExistAccountByUserID(a.UserID)
}

func (a *Auth) Add() error {
	return models.AddAccount(
		a.UserID,
		a.Username,
		a.LoginType,
		a.Country,
		a.Email,
		a.UserPicture)
}
