package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"

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
		part.Set("Content-Type", "image/png")
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
	}
	if err := writer.Close(); err != nil {
		log.Println("Failed To Close Writer")

		return nil, err
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
	buf := new(bytes.Buffer)

	if err := png.Encode(buf, *msg.imgContent); err != nil {
		log.Println("Failed To Encode png File")

		return
	}

	r, err := bot.PostFile(cid, "img.png", buf.Bytes())
	if err != nil {
		log.Println("Faild To Post Image")

		return
	}
	defer r.Body.Close()
	buf.Reset()
	if _, err := buf.ReadFrom(r.Body); err != nil {
		log.Println("Failed To Read Response Body")

		return
	}
	log.Printf("Status Code: %d\n", r.StatusCode)
	var f traq.FileInfo
	if err := json.Unmarshal(buf.Bytes(), &f); err != nil {
		log.Println("Failed To Decode JSON")

		return
	}

	{
		// log response
		log.Printf("File Info: %v\n", f)
		buf := new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(f); err != nil {
			log.Println("Failed To Encode JSON")

			return
		}
		log.Printf("Response Body: %s\n", buf.String())
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
