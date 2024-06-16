package bot

import (
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/traPtitech/go-traq"
)

type Message struct {
	strContent *string
	imgContent *image.Image
}

func NewMessage(strContent *string, imgContent *image.Image) *Message {
	return &Message{strContent: strContent, imgContent: imgContent}
}

func (bot *Bot) SendMessage(cid string, msg *Message, embed bool) {
	req := traq.NewPostMessageRequest(*msg.strContent)
	req.Embed = &embed
	m, r, err := bot.client.MessageApi.
		PostMessage(bot.auth, cid).
		PostMessageRequest(*req).
		Execute()
	if err != nil {
		log.Println("Failed To Post Text Message")

		return
	}
	log.Println("Sent Message: " + m.Content)
	log.Printf("StatusCode: %d\n", r.StatusCode)

	img, err := os.CreateTemp("", "img.jpeg")
	if err != nil {
		log.Println("Failed To Create Temp File")

		return
	}
	jpeg.Encode(img, *msg.imgContent, &jpeg.Options{Quality: 100})
	defer img.Close()

	f, r, err := bot.client.FileApi.
		PostFile(bot.auth).
		File(img).
		ChannelId(cid).
		Execute()
	if err != nil {
		log.Println("Faild To Post Image")

		return
	}
	log.Println("Sent File: " + f.GetName())
	log.Printf("Status Code: %d", r.StatusCode)
}
