package main

import (
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

	e := echo.New()

	bot := bot.New(
		os.Getenv("BOT_ID"),
		os.Getenv("BOT_USER_ID"),
		os.Getenv("BOT_ACCESS_TOKEN"),
		os.Getenv("VERIFICATION_TOKEN"),
	)

	macSecret, ok := os.LookupEnv("MAC_SECRET")
	if !ok {
		log.Fatal("MAC_SECRET is not set")
	}
	raspiSecret, ok := os.LookupEnv("RASPI_SECRET")
	if !ok {
		log.Fatal("RASPI_SECRET is not set")
	}
	sendChannelID, ok := os.LookupEnv("SEND_CHANNEL_ID")
	if !ok {
		log.Fatal("SEND_CHANNEL_ID is not set")
	}

	state := routes.NewState(bot, repo, macSecret, raspiSecret, sendChannelID)
	state.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
