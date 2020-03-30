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

//GetPostList get All PostList
func (p *PostController) GetPostList(c echo.Context) error {

	queryString := "SELECT p.id , s.subject_name as tag , p.tag_id , p.topic,p.location, a.username as tutor ,tic.amount,p.max_participant,p.start_at,p.end_at,p.price    from post p inner join account a on p.account_id = a.id inner join subject s on p.tag_id = s.tag_id INNER JOIN (SELECT post_id,count(post_id) as amount from ticket group by post_id) tic on p.id = tic.post_id"
	rows, err := p.DB.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	postList := make([]*model.PostList, 0)
	for rows.Next() {

		post := new(model.PostList)

		err := rows.Scan(&post.ID, &post.Tag, &post.TagID, &post.Topic, &post.Location, &post.Tutor, &post.Amount, &post.Full, &post.StartTime, &post.StopTime, &post.Price)
		if err != nil {
			return err
		}

		startDate := []rune(post.StartTime)
		stopDate := []rune(post.StopTime)

		date := string(startDate[0:10])
		length := len(startDate)
		var startTime = string(startDate[11:length])
		var stopTime = string(stopDate[11:length])
		post.Date = date
		post.StartTime = startTime
		post.StopTime = stopTime

		postList = append(postList, post)
	}

	return c.JSON(http.StatusOK, postList)
}
