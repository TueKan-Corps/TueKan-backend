package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

//Post routes for Post API
func Post(app *echo.Echo) {
	postController := controller.NewPostController(db.DB)

	postRoute := app.Group("/post")
	postRoute.POST("/", postController.CreatePost)
	postRoute.GET("/show", postController.GetAllPostByLimit)
}
