package routes

import (
	"github.com/labstack/echo/v4"
)

type GetTokenResponse struct {
	Token string `json:"token"`
}

// tokensテーブルをこのエンドポイント以外からmutateしてはいけない
// GET /stand
func (s *State) GetTokenHandler(c echo.Context) error {
	if s.macSecret != c.Request().Header.Get("X-Mac-Secret") {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	token, err := s.repo.ReadLatestToken()
	if err != nil {
		c.Logger().Error(err)

		return echo.NewHTTPError(500, "Internal server error")
	}

	return c.JSON(200, GetTokenResponse{token.Token})
}
