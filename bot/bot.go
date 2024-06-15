package bot

import (
	"context"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

type Bot struct {
	client            *traq.APIClient
	auth              context.Context
	Id                string
	UserId            string
	VerificationToken string
	AccessToken       string
}

func New(botId string, userId string, accessToken string, verificationToken string) Bot {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, accessToken)
	return Bot{
		client,
		auth,
		botId,
		userId,
		verificationToken,
		accessToken,
	}
}

func (bot *Bot) MakeHandlers() traqbot.EventHandlers {
	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(bot.PingHandler)
	return handlers
}
