package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// POST /checkin?token=foobar
func (s *State) CheckinHandler(c echo.Context) error {
	req := c.Request()
	traQID := req.Header.Get("X-Forwarded-User")

	if traQID == "" {
		return echo.NewHTTPError(401, "Unauthorized")
	}
	token := c.QueryParam("token")
	if token == "" {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	exists, err := s.repo.CheckIfTokenExists(token)
	if err != nil {
		c.Logger().Error(err)

		return echo.NewHTTPError(500, "Internal server error")
	}
	if !exists {
		return c.String(http.StatusForbidden, "Please recieve the QR token again")
	}

	s.raspiUser = token
	log.Println("POST /checkin id: " + traQID + ", token: " + token)

	return c.String(http.StatusOK, "POST /checkin id: "+traQID+", token: "+token)
}
