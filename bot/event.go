package bot

import (
	"log"

	traqbot "github.com/traPtitech/traq-bot"
)

func (bot *Bot) PingHandler(*traqbot.PingPayload) {
	log.Println("traQBot-pong")
}
