package routes

import (
	"TueKan-backend/controller"

	"github.com/labstack/echo"
)

// Home the beginning of greatness
func Home(app *echo.Echo) {
	app.GET("/", controller.Welcome)
}
