package routes

import (
	"TueKan-backend/controller"
	"TueKan-backend/db"

	"github.com/labstack/echo"
)

// Account routes for account
func Account(app *echo.Echo) {
	accountController := controller.NewAccountController(db.DB)

	accountRoute := app.Group("/account")
	accountRoute.POST("/", accountController.Create)
	accountRoute.POST("/img/:id", accountController.UploadProfileIMG)
	accountRoute.POST("/update", accountController.UpdateAccount)
	accountRoute.POST("/coin", accountController.UpdateCoin)

	accountRoute.GET("/", accountController.GetAll)
	accountRoute.GET("/img", accountController.GetProfileIMGList)
	accountRoute.GET("/img/:id", accountController.GetProfileIMG)
	accountRoute.GET("/:id", accountController.GetAccountById)

	accountRoute.DELETE("/img/cache", accountController.ClearIMGCache)
}
