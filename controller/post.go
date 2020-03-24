package controller

import (
	"TueKan-backend/model"
	"database/sql"
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
		return err
	}
	post.AccountID = accountID

	post.MaxParticipant, err = strconv.Atoi(c.FormValue("max_participant"))
	if err != nil {
		return err
	}

	post.TagID, err = strconv.Atoi(c.FormValue("tag_id"))
	if err != nil {
		return err
	}

	post.HeldAt = c.FormValue("held_at")

	currentTime := time.Now().Format("01-02-2006 15:04:05 Monday")

	queryString := "INSERT INTO post(account_id,topic,location,description,updated_at,created_at,held_at,tag_id,max_participant) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"

	_, err = p.DB.Exec(queryString, post.AccountID, post.Topic, post.Location, post.Description, currentTime, currentTime, post.HeldAt, post.TagID, post.MaxParticipant)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "Post created")
}

//GetAllPostByLimit get all post from db limit db params
func (p *PostController) GetAllPostByLimit(c echo.Context) error {

	limit := c.QueryParam("limit")

	queryString := "SELECT p.id,p.account_id,p.topic,p.location,p.description,p.updated_at,p.created_at,p.held_at,p.tag_id,p.max_participant,s.subject_name FROM post p LEFT JOIN  subject s ON p.tag_id = s.tag_id ORDER BY created_at DESC LIMIT $1"

	rows, err := p.DB.Query(queryString, limit)
	if err != nil {
		return err
	}
	defer rows.Close()

	posts := make([]*model.Post, 0)
	for rows.Next() {
		post := new(model.Post)

		err := rows.Scan(&post.ID, &post.AccountID, &post.Topic, &post.Location, &post.Description, &post.UpdatedAt, &post.CreatedAt, &post.HeldAt, &post.TagID, &post.MaxParticipant, &post.SubjectName)
		if err != nil {
			return err
		}

		posts = append(posts, post)
	}

	return c.JSON(http.StatusOK, posts)
}
