package routes

import (
	"TueKan-backend/controller"

	"github.com/labstack/echo"
)

// Index all routes
func Home(app *echo.Echo) {
	app.GET("/", controller.Welcome)
}
