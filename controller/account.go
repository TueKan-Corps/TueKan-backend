package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
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

	queryString := "SELECT id,username,password,coin_amount,first_name,last_name FROM account ORDER BY id"

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
			fmt.Println(err)
			return err
		}


		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (a *AccountController) UploadProfileIMG(c echo.Context) error {

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// receive file from user
	file, err :=  c.FormFile("profile_img")
	if err != nil{
		return err
	}
	src, err := file.Open()
	if err != nil{
		return err
	}
	defer src.Close()

	// Destination
	imgPath :=  "./img/"+fmt.Sprintf("%d",accountID)+".jpg"
	dst, err := os.Create(imgPath)
	if err!= nil{
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// Save the file path in db
	queryString := "UPDATE account SET profile_img_path=$1 WHERE id=$2"
	_, err = a.DB.Exec(queryString,imgPath,accountID)
	if err != nil{
		return err
	}

	return c.JSON(http.StatusOK, "Profile image uploaded")
}
