package routes

import (
	"TueKan-backend/controller"

	"github.com/labstack/echo"
)

// Index all routes
func Index(app *echo.Echo) {
	app.GET("/", controller.Welcome)
}
