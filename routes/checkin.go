package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// POST /checkin?token=foobar
func (s *State) CheckinHandler(c echo.Context) error {
	req := c.Request()
	userID := req.Header.Get("X-Forwarded-User")

	if userID == "" {
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

	s.raspiUser = userID
	log.Println("POST /checkin id: " + userID + ", token: " + token)

	type jsonID struct {
		UserID string `json:"userID,omitempty"`
	}

	return c.JSON(http.StatusOK, &jsonID{UserID: userID})
}
