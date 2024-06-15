package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *State) PostStandHandler(c echo.Context) error {

	return c.String(http.StatusOK, "POST stand")
}
