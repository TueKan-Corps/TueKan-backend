package model

import (
	"golang.org/x/crypto/bcrypt"
)

// Contact type
type Contact struct {
	ID   int
	Name string
	Link string
}

// Account model
type Account struct {
	ID             int        `json:"id" db:"id"`
	Username       string     `json:"username" db:"username"`
	Password       string     `json:"password" db:"password"`
	FirstName      string     `json:"first_name" db:"first_name"`
	LastName       string     `json:"last_name" db:"last_name"`
	CoinAmount     int        `json:"coin_amount" db:"coin_amount"`
	ProfileIMGPath string     `json:"profile_img_path,omitempty" db:"profile_img_path,omitempty"`
	Contact        [5]Contact `json:"contact" db:"contact"`
}

func (a *Account) HashAndSaltPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	a.Password = string(hash)
	return nil
}
