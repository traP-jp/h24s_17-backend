package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello/:name", func(c echo.Context) error {
		userID := c.Param("name")

		return c.String(http.StatusOK, "Hello, "+userID+".\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
