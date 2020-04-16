package main

import (
	"TueKan-backend/config"
	"TueKan-backend/db"
	"TueKan-backend/routes"
	"TueKan-backend/thirdparty"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

func main() {
	var c config.Config
	var err error

	// Load secret from .env file
	if err := c.Init(); err != nil {
		log.Fatal("Load .env failed", err)
	}

	// Connect to DB
	if err := db.Init(&c); err != nil {
		log.Fatal("Create a connection to db failed", err)
	}

	// Connect to AWS
	err := thirdparty.InitAWSSession(&c)
	if err != nil {
		log.Fatal("connect to AWS failed", err)
	}

	app := echo.New()
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.Home(app)
	routes.Account(app)
	routes.Post(app)
	routes.Subject(app)
	routes.Auth(app)
	routes.Session(app)

	var port = fmt.Sprintf(":%s", c.Port)

	app.Logger.Fatal(app.Start(port))
}
