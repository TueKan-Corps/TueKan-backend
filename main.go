package main

import (
	"TueKan-backend/config"
	"TueKan-backend/db"
	"TueKan-backend/routes"
	"TueKan-backend/thirdparty"
	"TueKan-backend/util"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"time"
)

func main() {
	var c config.Config

	// Load secret from .env file
	if err := c.Init(); err != nil {
		log.Fatal("Load .env failed", err)
	}

	// Connect to DB
	if err := db.Init(&c); err != nil {
		log.Fatal("Create a connection to db failed", err)
	}

	// Connect to AWS
	if err := thirdparty.InitAWSSession(&c); err != nil {
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

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for t := range ticker.C {
			_ = t // we don't print the ticker time, so assign this `t` variable to underscore `_` to avoid error
			err := util.ClearOutdatedPost(db.DB)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("outdated posts cleared")
			}
		}
	}()

	app.Logger.Fatal(app.Start(port))
}
