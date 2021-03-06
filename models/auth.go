package models

import "github.com/jinzhu/gorm"

type Account struct {
	ID          int    `gorm:"primary_key" json:"id"`
	UserID      string `json:"username"`
	Username    string `json:"user_id"`
	LoginType   string `json:"login_type"`
	Country     string `json:"country"`
	Email       string `json:"email"`
	UserPicture string `json:"user_picture"`
	TeamID      string `json:"team_id"`
}

// ExistAccountByUserID checks if there is a account with the same user_id
func ExistAccountByUserID(userid string) (*Account, error) {
	var account Account
	err := db.First(&account, "user_id = ?", userid).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if account.ID > 0 {
		return &account, nil
	}

	return nil, nil
}

func AddAccount(userId, name, loginType, country, email, picture string) error {
	account := Account{
		UserID:      userId,
		Username:    name,
		LoginType:   loginType,
		Country:     country,
		Email:       email,
		UserPicture: picture,
	}
	if err := db.Create(&account).Error; err != nil {
		return err
	}
	return nil
}

func GetAccount() {

}
