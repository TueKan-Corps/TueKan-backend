package controller

import (
	"TueKan-backend/model"
	"TueKan-backend/thirdparty"
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

	// hash password
	if err := account.HashAndSaltPassword(); err != nil {
		return err
	}

	coinAmount, err := strconv.Atoi(c.FormValue("coin_amount"))
	if err != nil {
		return err
	}
	account.CoinAmount = coinAmount

	account.FirstName = c.FormValue("first_name")
	account.LastName = c.FormValue("last_name")

	account.Contact = getContactFromContext(c)

	// insert general info
	queryString := `INSERT INTO account (username,password,coin_amount,first_name,last_name,contact) 
					VALUES ($1,$2,$3,$4,$5,
					ARRAY[$6,$7,$8,$9,$10]) RETURNING id`
	err = a.DB.QueryRow(queryString,
		account.Username,
		account.Password,
		account.CoinAmount,
		account.FirstName,
		account.LastName,
		account.Contact[0].Link,
		account.Contact[1].Link,
		account.Contact[2].Link,
		account.Contact[3].Link,
		account.Contact[4].Link).Scan(&account.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, account.ID)
}

// GetAll get all account
func (a *AccountController) GetAll(c echo.Context) error {

	queryString := "SELECT id,username,password,coin_amount,first_name,last_name,description FROM account ORDER BY id"

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
			&account.LastName,
			&account.Description)

		account.Contact, err = getContactFromDB(a.DB, account.ID)
		if err != nil {
			return err
		}

		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

// UploadProfileIMG upload a profile image
func (a *AccountController) UploadProfileIMG(c echo.Context) error {

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	// receive file from user
	file, err := c.FormFile("profile_img")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	filename := fmt.Sprintf("%d.jpg", accountID)

	// Destination
	imgPath := "./img/" + filename
	dst, err := os.Create(imgPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	err = thirdparty.UploadItem(filename)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "couldn't upload")
	}

	return c.JSON(http.StatusOK, "Profile image uploaded")
}

func (a *AccountController) GetProfileIMGList(c echo.Context) error {
	fileItems, err := thirdparty.ListItems()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, fileItems)
}

// GetProfileIMG get account profile image
func (a *AccountController) GetProfileIMG(c echo.Context) error {
	accountID := c.Param("id")

	filename := accountID + ".jpg"

	err := thirdparty.DownloadItem(filename)
	if err != nil {
		c.JSON(http.StatusNotFound, "cannot find img")
	}

	filepath := "./img/" + filename

	return c.File(filepath)
}

func (a *AccountController) ClearIMGCache(c echo.Context) error {
	err := os.RemoveAll("img")
	if err != nil {
		return err
	}
	err = os.MkdirAll("img", 0777)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "cleared")
}

func getContactFromContext(c echo.Context) [5]model.Contact {
	var contacts [5]model.Contact
	contactList := []string{"facebook", "instagram", "youtube", "email", "website"}

	for i := 1; i < 6; i++ {
		contacts[i-1] = model.Contact{
			ID:   i,
			Name: contactList[i-1],
			Link: c.FormValue(contactList[i-1]),
		}
	}

	return contacts
}

func getContactFromDB(db *sql.DB, id int) ([5]model.Contact, error) {
	var contacts [5]model.Contact
	contactList := []string{"facebook", "instagram", "youtube", "email", "website"}

	queryString := "SELECT contact[$1] FROM account WHERE id=$2"

	for i := 1; i < 6; i++ {

		var link string

		row := db.QueryRow(queryString, i, id)
		err := row.Scan(&link)
		if err != nil {
			return [5]model.Contact{}, err
		}

		contacts[i-1] = model.Contact{
			ID:   i,
			Name: contactList[i-1],
			Link: link,
		}
	}

	return contacts, nil
}

func (a *AccountController) GetAccountById(c echo.Context) error {

	accountID, err := strconv.Atoi(c.Param("id"))
	queryString := "SELECT id,username,password,coin_amount,first_name,last_name,description FROM account WHERE id = $1"

	rows, err := a.DB.Query(queryString, accountID)
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
			&account.LastName,
			&account.Description)

		account.Contact, err = getContactFromDB(a.DB, account.ID)
		if err != nil {
			return err
		}

		accounts = append(accounts, account)
	}

	return c.JSON(http.StatusOK, accounts)
}

func (a *AccountController) UpdateAccount(c echo.Context) error {
	account := new(model.UpdateAccount)

	if err := c.Bind(account); err != nil {
		return err
	}

	queryString := "UPDATE account SET  first_name = $1 , last_name = $2,description = $3, contact = ARRAY[$4,$5,$6,$7,$8]WHERE id = $9"
	_, err := a.DB.Exec(queryString,
		account.FirstName,
		account.LastName,
		account.Description,
		account.Contact[0].Link,
		account.Contact[1].Link,
		account.Contact[2].Link,
		account.Contact[3].Link,
		account.Contact[4].Link,
		account.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Account update")
}

func (a *AccountController) UpdateCoin(c echo.Context) error {

	isOverwrite := c.FormValue("is_overwrite")

	coin := new(model.UpdateCoins)
	if err := c.Bind(coin); err != nil {
		return err
	}

	var queryString string
	if isOverwrite == "true" {
		queryString = "UPDATE account SET  coin_amount = $1 WHERE id = $2"
	} else {
		queryString = "UPDATE account SET coin_amount = coin_amount+$1 WHERE id = $2"
	}

	_, err := a.DB.Exec(queryString, coin.Coin, coin.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, "Coin update")
}
