package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"
	"github.com/labstack/echo"
)

func Subject(app *echo.Echo)  {
	subjectController := controller.NewSubjectController(db.DB)

	subjectRoute := app.Group("/subject")
	subjectRoute.GET("/",subjectController.GetAllSubject)
	subjectRoute.POST("/",subjectController.CreateNewSubject)
}