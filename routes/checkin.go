package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *State) CheckinHandler(c echo.Context) error {
	req := c.Request()
	traQID := req.Header.Get("X-Forwarded-User")

	log.Println("POST /checkin from " + traQID)

	return c.String(http.StatusOK, "POST /checkin from "+traQID)
}
