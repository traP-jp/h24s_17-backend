package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")

		return c.String(http.StatusOK, "Hello, "+name+".\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
