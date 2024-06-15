package bot

import (
	"fmt"

	traqbot "github.com/traPtitech/traq-bot"
)

func (bot *Bot) PingHandler(payload *traqbot.PingPayload) {
	fmt.Println("pong")
}
