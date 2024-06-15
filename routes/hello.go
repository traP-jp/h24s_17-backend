package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *State) HelloHandler(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, "Hello, "+name+".\n")
}
