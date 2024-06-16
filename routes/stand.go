package routes

import (
	"image/png"
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
	if !strings.HasPrefix(req.Header.Get("Content-Type"), "image/png") {
		return echo.NewHTTPError(415, "Unsupported Media Type")
	}
	image, err := png.Decode(req.Body)
	if err != nil {
		return echo.NewHTTPError(400, "Bad Request")
	}
	s.bot.SendImage(
		s.sendChannelID,
		bot.NewMessage(&image),
		true)

	return c.String(http.StatusOK, "POST /state")
}
