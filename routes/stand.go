package routes

import (
	"image/jpeg"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/bot"
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

	image, err := jpeg.Decode(req.Body)
	if err != nil {
		return echo.NewHTTPError(400, "Bad Request")
	}
	s.bot.SendMessage(
		s.sendChannelID,
		bot.NewMessage(&traQID, &image),
		true)

	return c.String(http.StatusOK, "POST /state")
}
