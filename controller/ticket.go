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

	queryString := "SELECT p.id, p.account_id, s.subject_name as tag, p.tag_id, p.topic, p.location, a.username     as tutor, tic.count      as amount, p.max_participant as max , p.start_at , p.end_at, p.price, t.access_code, p.description from post p INNER JOIN (SELECT post_id, count(post_id) as count from ticket group by post_id) tic on p.id = tic.post_id INNER JOIN subject s on p.tag_id = s.tag_id INNER JOIN account a on p.account_id = a.id INNER JOIN ticket t on p.id = t.post_id where t.account_id = $1 order by start_at asc"
	rows, err := t.DB.Query(queryString, accountID)
	if err != nil {
		return err
	}
	defer rows.Close()

	ticketList := make([]*model.TicketList, 0)
	for rows.Next() {

		ticket := new(model.TicketList)

		err := rows.Scan(&ticket.ID, &ticket.AccountID, &ticket.Tag, &ticket.TagID, &ticket.Topic, &ticket.Location, &ticket.Tutor, &ticket.Amount, &ticket.Full, &ticket.StartTime, &ticket.StopTime, &ticket.Price, &ticket.Ticket, &ticket.Description)
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

func (t *TicketController) Redeem(c echo.Context) error {
	ticketRedeem := new(model.RedeemTicket)
	if err := c.Bind(ticketRedeem); err != nil {
		return err
	}

	accessCode := strconv.Itoa(ticketRedeem.AccessCode)

	queryString := "UPDATE ticket set isredeem = $1 where post_id = $2 and access_code = $3"
	_, err := t.DB.Query(queryString, ticketRedeem.IsRedeem, ticketRedeem.PostID, accessCode)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "Redeemed")

}

func (t *TicketController) CreateTicket(c echo.Context) error {
	Ticket := new(model.NewTicket)
	if err := c.Bind(Ticket); err != nil {
		return err
	}

	queryString := `INSERT INTO ticket(account_id, post_id, isredeem, access_code)
					values ($1, $2, false, (SELECT num
                       FROM GENERATE_SERIES(1000, 9999) AS s(num)
                       order by random()
                       LIMIT 1));`
	_, err := t.DB.Query(queryString, Ticket.AccountID, Ticket.PostID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, "Create Ticket")

}
