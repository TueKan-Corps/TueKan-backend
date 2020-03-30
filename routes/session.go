package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"
	"github.com/labstack/echo"
)

//Session routes for Post API
func Session(app *echo.Echo) {
	sessionController := controller.NewSessionController(db.DB)

	sessionRoute := app.Group("/session")
	sessionRoute.GET("/", sessionController.GetAll)
	sessionRoute.DELETE("/", sessionController.ClearAll)
}
