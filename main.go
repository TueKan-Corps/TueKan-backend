package main

import (
	"TueKan-backend/config"
	"TueKan-backend/db"
	"TueKan-backend/routes"
	"TueKan-backend/util"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// Connect to AWS
	err = util.InitAWSSession(&c)
	if err != nil {
		log.Fatal("connect to AWS failed", err)
	}

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.Home(app)
	routes.Account(app)
	routes.Post(app)
	routes.Subject(app)
	routes.Ticket(app)
	routes.Auth(app)
	routes.Session(app)

	var port = fmt.Sprintf(":%s", c.Port)

	app.Logger.Fatal(app.Start(port))
}
