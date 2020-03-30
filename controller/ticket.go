package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type TicketController struct {
	DB *sql.DB
}

//NewTicketController create  Post controller
func NewTicketController(db *sql.DB) *TicketController {
	return &TicketController{DB: db}
}

func (t *TicketController) GetTicket(c echo.Context) error {

	accountID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	queryString := "SELECT p.id, s.subject_name as tag, p.tag_id, p.topic, p.location, a.username     as tutor, tic.count      as amount, p.max_participant as max , p.start_at , p.end_at, p.price, t.access_code, p.description from post p INNER JOIN (SELECT post_id, count(post_id) as count from ticket group by post_id) tic on p.id = tic.post_id INNER JOIN subject s on p.tag_id = s.tag_id INNER JOIN account a on p.account_id = a.id INNER JOIN ticket t on p.id = t.post_id where t.account_id = $1"
	rows, err := t.DB.Query(queryString, accountID)
	if err != nil {
		return err
	}
	defer rows.Close()

	ticketList := make([]*model.TicketList, 0)
	for rows.Next() {

		ticket := new(model.TicketList)

		err := rows.Scan(&ticket.ID, &ticket.Tag, &ticket.TagID, &ticket.Topic, &ticket.Location, &ticket.Tutor, &ticket.Amount, &ticket.Full, &ticket.StartTime, &ticket.StopTime, &ticket.Price, &ticket.Ticket, &ticket.Description)
		if err != nil {
			return err
		}

		startDate := []rune(ticket.StartTime)
		stopDate := []rune(ticket.StopTime)

		date := string(startDate[0:10])
		length := len(startDate)
		var startTime = string(startDate[11:length])
		var stopTime = string(stopDate[11:length])
		ticket.Date = date
		ticket.StartTime = startTime
		ticket.StopTime = stopTime

		ticketList = append(ticketList, ticket)
	}

	return c.JSON(http.StatusOK, ticketList)

}
