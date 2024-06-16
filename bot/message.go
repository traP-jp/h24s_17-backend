package bot

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/traPtitech/go-traq"
)

type Message struct {
	imgContent *image.Image
}

func NewMessage(imgContent *image.Image) *Message {
	return &Message{imgContent: imgContent}
}

func (bot *Bot) SendImage(cid string, msg *Message, embed bool) {
	img, err := os.CreateTemp("", "img.jpeg")
	if err != nil {
		log.Println("Failed To Create Temp File")

		return
	}
	defer os.Remove(img.Name())
	err = jpeg.Encode(img, *msg.imgContent, nil)

	if err != nil {
		log.Println("Failed To Encode jpeg File")

		return
	}

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

	if err := img.Close(); err != nil {
		log.Println("Failed To Close File")
	}

	file := fmt.Sprintf("https://q.trap.jp/files/%s", f.Id)
	req := traq.NewPostMessageRequest(file)
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

}
