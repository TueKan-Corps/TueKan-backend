package main

import (
	"TueKan-backend/config"
	"TueKan-backend/db"
	"TueKan-backend/routes"
	"log"

	"github.com/labstack/echo"
)

func main() {
	var config config.Config

	// Load secret from .env file
	err := config.Init()
	if err != nil {
		log.Fatal("Load .env failed", err)
	}

	// Connect to DB
	err = db.Init(&config)
	if err != nil {
		log.Fatal("Create a connection to db failed", err)
	}

	app := echo.New()

	routes.Index(app)

	app.Logger.Fatal(app.Start(":1323"))
}
