package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"net/http"

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
	post := new(model.CreatePost)

	if err := c.Bind(post); err != nil {
		return err
	}

	queryString := "INSERT INTO post(account_id,topic,location,description,updated_at,created_at) VALUES($1,$2,$3,$4,$5,$6)"
	_, err := p.DB.Exec(queryString, post.AccountID, post.Topic, post.Location, post.Description, post.UpdatedAt, post.CreatedAt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "Account created")
}
