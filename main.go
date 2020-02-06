package main

import (
	"TueKan-backend/routes"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()

	routes.Index(app)

	app.Logger.Fatal(app.Start(":1323"))
}
