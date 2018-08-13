package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Test お試しAPI
func Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}
