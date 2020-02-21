package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

// Account routes for account
func Account(app *echo.Echo) {
	accountController := controller.NewAccountController(db.DB)

	accountRoute := app.Group("/account")
	accountRoute.POST("/", accountController.Create)
	accountRoute.GET("/", accountController.GetAll)
}
