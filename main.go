package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/bot"
	"github.com/traP-jp/h24s_17-backend/models"
	"github.com/traP-jp/h24s_17-backend/routes"
)

func main() {
	repo, err := models.ConnectRepository()
	if err != nil {
		log.Fatal(err)

		return
	}
	err = repo.Migrate()
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

	bot := bot.New(
		os.Getenv("BOT_ID"),
		os.Getenv("BOT_USER_ID"),
		os.Getenv("BOT_ACCESS_TOKEN"),
		os.Getenv("VERIFICATION_TOKEN"),
	)
	state := routes.NewState(bot, repo)
	state.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
