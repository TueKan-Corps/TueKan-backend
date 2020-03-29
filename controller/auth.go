package controller

import (
	"TueKan-backend/model"
	"github.com/labstack/echo"
	"net/http"
)

func (a *AccountController) Login(c echo.Context) error {
	inputUsername := c.FormValue("username")
	inputPassword := c.FormValue("password")

	account := new(model.Account)

	queryString := "SELECT password FROM account WHERE username=$1"
	row := a.DB.QueryRow(queryString, inputUsername)
	err := row.Scan(&account.Password)
	if err != nil {
		return err
	}
	account.Username = inputUsername

	if err := account.ComparePassword(inputPassword); err != nil {
		return err
	}

	// return these things as a key
	key := new(model.Key)

	queryString = "SELECT id FROM  account WHERE username=$1"
	row = a.DB.QueryRow(queryString, inputUsername)
	err = row.Scan(&key.AccountID)
	if err != nil {
		return err
	}
	key.SecretKey = "very_secret"
	return c.JSON(http.StatusOK, key)
}
