package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// AccountController a controller of an account model
type AccountController struct {
	DB *sql.DB
}

// NewAccountController create a new account controller
func NewAccountController(db *sql.DB) *AccountController {
	return &AccountController{DB: db}
}

// Create add new account
func (a *AccountController) Create(c echo.Context) error {
	account := new(model.Account)

	if err := c.Bind(account); err != nil {
		return err
	}

	coinAmount, err := strconv.Atoi(c.FormValue("coin_amount"))
	if err != nil {
		return err
	}
	account.CoinAmount = coinAmount

	queryString := "INSERT INTO account (username,password,coin_amount,first_name,last_name) VALUES ($1,$2,$3,$4,$5)"
	_, err = a.DB.Exec(queryString,
		account.Username,
		account.Password,
		account.CoinAmount,
		account.FirstName,
		account.LastName)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Account created")
}

// GetAll get all account
func (a *AccountController) GetAll(c echo.Context) error {

	queryString := "SELECT * FROM account ORDER BY id"
	rows, err := a.DB.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	accounts := make([]*model.Account, 0)
	for rows.Next() {
		account := new(model.Account)

		err := rows.Scan(&account.ID,
			&account.Username,
			&account.Password,
			&account.CoinAmount,
			&account.FirstName,
			&account.LastName)

		if err != nil {
			return err
		}
		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}
