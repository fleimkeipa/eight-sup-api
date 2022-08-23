package utils

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Authorization() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(s1, s2 string, c echo.Context) (bool, error) {
		if s1 == os.Getenv("username") && s2 == os.Getenv("password") {
			return true, nil
		} else {
			return false, nil
		}
	})
}
