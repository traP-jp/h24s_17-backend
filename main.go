package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/models"
	"github.com/traP-jp/h24s_17-backend/routes"
)

func main() {
	_, err := models.Connect()
	if err != nil {
		log.Fatal(err)

		return
	}

	macSecret, ok := os.LookupEnv("MAC_SECRET")
	if !ok {
		fmt.Println("MAC_SECRET is not set")

		return
	}
	fmt.Printf("loaded MAC_SECRET: %s\n", macSecret)

	e := echo.New()
	routes.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
