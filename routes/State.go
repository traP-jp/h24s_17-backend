package routes

import (
	"github.com/traP-jp/h24s_17-backend/bot"
	"github.com/traP-jp/h24s_17-backend/models"
)

type State struct {
	bot           *bot.Bot
	repo          *models.Repository
	macSecret     string
	raspiSecret   string
	sendChannelID string
	raspiUser     string
}

func NewState(b *bot.Bot, r *models.Repository, macSecret string, raspiSecret string, sendChannelID string) *State {
	return &State{bot: b, repo: r, macSecret: macSecret, raspiSecret: raspiSecret, sendChannelID: sendChannelID, raspiUser: ""}
}
