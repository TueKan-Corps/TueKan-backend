package controller

import (
	"TueKan-backend/model"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"strings"
	"time"
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
	session := new(model.Session)

	queryString = "SELECT id FROM  account WHERE username=$1"
	row = a.DB.QueryRow(queryString, inputUsername)
	err = row.Scan(&session.AccountID)
	if err != nil {
		return err
	}

	// Generate random token
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 100
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	session.Token = str

	// Record a token into Session
	queryString = "INSERT INTO session(account_id, token) VALUES ($1,$2) RETURNING id"

	row = a.DB.QueryRow(queryString, session.AccountID, session.Token)
	err = row.Scan(&session.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, session)
}
