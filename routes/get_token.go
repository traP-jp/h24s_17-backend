package routes

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type GetTokenResponse struct {
	Token string `json:"token"`
}

// tokensテーブルをこのエンドポイント以外からmutateしてはいけない
// GET /stand
func (s *State) GetTokenHandler(c echo.Context) error {
	// mac_secretが環境変数と一致
	macSecret, ok := os.LookupEnv("MAC_SECRET")
	if !ok {
		fmt.Println("MAC_SECRET is not set")
		return echo.NewHTTPError(500, "Internal server error")
	}

	if macSecret != c.Request().Header.Get("X-Mac-Secret") {
		return echo.NewHTTPError(401, "Unauthorized")
	}

	token, err := s.repo.ReadLatestToken()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(500, "Internal server error")
	}
	return c.JSON(200, GetTokenResponse{token.Token})
}
