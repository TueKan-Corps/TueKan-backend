package util

import "github.com/labstack/echo"

func CheckToken(key string, c echo.Context) (bool, error) {
	return len(key) == 100, nil
}
