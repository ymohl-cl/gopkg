package server

import (
	"net/http"

	"github.com/labstack/echo"
)

// Ping method http GET
func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK, &ModelPong{Pong: true})
}
