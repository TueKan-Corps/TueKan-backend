package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

//PostController for Post model
type PostController struct {
	DB *sql.DB
}

//NewPostController create  Post controller
func NewPostController(db *sql.DB) *PostController {
	return &PostController{DB: db}
}

//CreatePost from json body
func (p *PostController) CreatePost(c echo.Context) error {
	post := new(model.Post)

	if err := c.Bind(post); err != nil {
		return err
	}

	accountID, err := strconv.Atoi(c.FormValue("account_id"))
	if err != nil {
		return nil
	}
	post.AccountID = accountID

	heldAt := c.FormValue("held_at")

	currentTime := time.Now().Format("01-02-2006 15:04:05 Monday")

	queryString := "INSERT INTO post(account_id,topic,location,description,updated_at,created_at,held_at,tag) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err = p.DB.Exec(queryString, post.AccountID, post.Topic, post.Location, post.Description, currentTime, currentTime, heldAt, post.Tag)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusCreated, "Post created")
}

//GetAllPostByLimit get all post from db limit db params
func (p *PostController) GetAllPostByLimit(c echo.Context) error {

	limit := c.QueryParam("limit")

	queryString := "SELECT * FROM post ORDER BY created_at DESC LIMIT $1"

	rows, err := p.DB.Query(queryString, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)

		err := rows.Scan(&post.ID, &post.AccountID, &post.Topic, &post.Location, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.HeldAt, &post.Tag)
		if err != nil {
			return err
		}
		posts = append(posts, post)
	}

	return c.JSON(http.StatusOK, posts)
}
