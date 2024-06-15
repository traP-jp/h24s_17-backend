package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	MAC_SECRET, ok := os.LookupEnv("MAC_SECRET")
	if !ok {
		fmt.Println("MAC_SECRET is not set")
		return
	}
	fmt.Printf("loaded MAC_SECRET: %s\n", MAC_SECRET)

	e := echo.New()

	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")

		return c.String(http.StatusOK, "Hello, "+name+".\n")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
