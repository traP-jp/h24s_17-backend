package bot

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/traPtitech/go-traq"
)

type Message struct {
	imgContent *image.Image
}

func NewMessage(imgContent *image.Image) *Message {
	return &Message{imgContent: imgContent}
}

func (bot *Bot) PostFile(cid string, filename string, content []byte) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	{
		part := make(textproto.MIMEHeader)
		part.Set("Content-Type", "image/jpeg")
		part.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))
		wp, err := writer.CreatePart(part)
		if err != nil {
			log.Println("Failed To Write Content")

			return nil, err
		}
		if _, err = wp.Write(content); err != nil {
			log.Println("Failed To Write Content")

			return nil, err
		}
	}
	{
		part, err := writer.CreateFormField("channelId")
		if err != nil {
			log.Println("Failed To Create Form Field")

			return nil, err
		}
		_, err = part.Write([]byte(cid))
		if err != nil {
			log.Println("Failed To Write Content")

			return nil, err
		}
		err = writer.Close()
		if err != nil {
			log.Println("Failed To Close Writer")

			return nil, err
		}
	}
	r, err := http.NewRequestWithContext(bot.auth, "POST", "https://q.trap.jp/api/v3/files", body)
	if err != nil {
		log.Println("Failed To Create Request")

		return nil, err
	}
	r.Header.Set("Content-Type", writer.FormDataContentType())
	r.Header.Set("Authorization", "Bearer "+bot.auth.Value(traq.ContextAccessToken).(string))
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		log.Println("Failed To Do Request")

		return nil, err
	}
	log.Printf("Status Code: %d\n", resp.StatusCode)

	return resp, nil
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
