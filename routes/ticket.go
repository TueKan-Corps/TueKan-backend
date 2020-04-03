package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

//Post routes for Post API
func Ticket(app *echo.Echo) {

	ticketController := controller.NewTicketController(db.DB)

	ticketRoute := app.Group("/ticket")

	ticketRoute.GET("/:id", ticketController.GetTicket)
	ticketRoute.POST("/redeem", ticketController.Redeem)
	ticketRoute.POST("/", ticketController.CreateTicket)
	//ticketRoute.GET("/users/:id", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "/users/:id")
	//})
}
