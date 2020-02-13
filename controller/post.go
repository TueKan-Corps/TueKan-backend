package controller

import (
	"database/sql"

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

	return nil
}
