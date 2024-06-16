package bot

import (
	"image"
	"log"

	"github.com/traPtitech/go-traq"
)

type Message struct {
	strContent *string
	imgContent *image.Image
}

func (bot *Bot) SendMessage(cid string, msg *Message, embed bool) {
	req := traq.NewPostMessageRequest(*msg.strContent)
	req.Embed = &embed
	m, r, err := bot.client.MessageApi.
		PostMessage(bot.auth, cid).
		PostMessageRequest(*req).
		Execute()
	if err != nil {
		bot.client.MessageApi.
			PostMessage(bot.auth, cid).
			PostMessageRequest(*traq.NewPostMessageRequest("Failed To Post Message")).
			Execute()
	}
	log.Println("Sent Message: " + m.Content)
	log.Printf("StatusCode: %d\n", r.StatusCode)
}
