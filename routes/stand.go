package routes

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s *State) PostStandHandler(c echo.Context) error {
	req := c.Request()
	if req.Header.Get("X-Raspi-Secret") != s.raspiSecret {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	if !strings.HasPrefix(req.Header.Get("Content-Type"), "image/jpeg") {
		return echo.NewHTTPError(415, "Unsupported Media Type")
	}
	traQID := req.Header.Get("X-Forwarded-User")

	// TODO

	return c.String(http.StatusOK, "POST /state")
}
