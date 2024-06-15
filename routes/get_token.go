package router

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/utils"
)

type GetTokenResponse struct {
	token string `json:token`
}

// tokensテーブルをこのエンドポイント以外からmutateしてはいけない
// GET /stand?mac_secret=foobar
func GetTokenhandler(c echo.Context, db sqlx.DB) error {
	// mac_secretが環境変数と一致

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
