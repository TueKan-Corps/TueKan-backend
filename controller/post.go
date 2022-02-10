package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"strconv"

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

//GetPosting get your post by account id
func (p *PostController) GetPosting(c echo.Context) error {

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	queryString := "SELECT p.id, s.subject_name as tag, p.tag_id, p.topic, p.location, a.username     as tutor, tic.cnt        as amount, p.max_participant as max, p.start_at, p.end_at, p.price, p.description, par.list FROM post p INNER JOIN subject s on P.tag_id = s.tag_id INNER JOIN account a on p.account_id = a.id INNER JOIN (SELECT post_id, count(post_id) AS cnt FROM ticket GROUP BY post_id) tic on p.id = tic.post_id INNER JOIN (SELECT post_id, json_agg(json_build_object('id', account_id, 'ticket', access_code, 'first_name', a2.first_name, 'last_name', a2.last_name)) as list FROM ticket INNER JOIN account a2 on ticket.account_id = a2.id group by post_id) par on par.post_id = p.id where p.account_id = $1"
	rows, err := p.DB.Query(queryString, accountID)
	if err != nil {
		return err
	}
	defer rows.Close()

	postingList := make([]*model.Posting, 0)

	for rows.Next() {

		posting := new(model.Posting)

		err := rows.Scan(&posting.ID, &posting.Tag, &posting.TagID, &posting.Topic, &posting.Location, &posting.Tutor, &posting.Amount, &posting.Full, &posting.StartTime, &posting.StopTime, &posting.Price, &posting.Description, &posting.Participant)
		if err != nil {
			return err
		}

		startDate := []rune(posting.StartTime)
		stopDate := []rune(posting.StopTime)

		date := string(startDate[0:10])
		length := len(startDate)
		var startTime = string(startDate[11:length])
		var stopTime = string(stopDate[11:length])
		posting.Date = date
		posting.StartTime = startTime
		posting.StopTime = stopTime

		postingList = append(postingList, posting)
	}

	return c.JSON(http.StatusOK, postingList)

}
