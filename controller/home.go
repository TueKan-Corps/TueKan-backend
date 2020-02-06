package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// Welcome return welcome string
func Welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to TueKan Backend")
}
