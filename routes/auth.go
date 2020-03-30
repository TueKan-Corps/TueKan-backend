package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

// Account routes for account
func Auth(app *echo.Echo) {
	accountController := controller.NewAccountController(db.DB)

	accountRoute := app.Group("/auth")
	accountRoute.POST("/login", accountController.Login)
	accountRoute.POST("/logout", accountController.Logout)
}
