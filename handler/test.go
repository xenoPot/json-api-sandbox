package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// Test お試しAPI
func Test() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonData := map[string]string{
			"message": "Hello World",
			"data":    "test",
		}
		return c.JSON(http.StatusOK, jsonData)
	}
}
