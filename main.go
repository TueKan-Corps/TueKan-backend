package main

import (
	"TueKan-backend/config"
	"TueKan-backend/db"
	"TueKan-backend/routes"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

func main() {
	var c config.Config

	// Load secret from .env file
	err := c.Init()
	if err != nil {
		log.Fatal("Load .env failed", err)
	}

	// Connect to DB
	err = db.Init(&c)
	if err != nil {
		log.Fatal("Create a connection to db failed", err)
	}

	app := echo.New()

	routes.Home(app)
	routes.Account(app)
	routes.Post(app)

	var port = fmt.Sprintf(":%s", c.Port)
	app.Logger.Fatal(app.Start(port))
}
