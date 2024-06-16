package routes

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/utils"
)

type GetTokenResponse struct {
	token string `json:token`
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

	// tokenをdbから取ってくる
	token := Token{}
	err := db.Get(&token, "SELECT * FROM tokens WHERE DATE_ADD(created_at, INTERVAL 5 MINUTE) <= NOW()")

	// 有効なtokenが存在しない場合 作成してinsertする
	if err != nil {
		newToken, err := utils.GenerateRandomToken(10)
		if err != nil {
			return echo.NewHTTPError(500, "Internal server error")

		}

		_, err = db.Exec("INSERT INTO tokens (token, created_at) VALUES (?, NOW())", newToken)

		if err != nil {
			return echo.NewHTTPError(500, "Internal server error")
		}
		return c.JSON(200, GetTokenResponse{token})
	}
	return c.JSON(200, GetTokenResponse{token})
}