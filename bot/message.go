package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/traPtitech/go-traq"
)

type Message struct {
	userID     string
	imgContent *image.Image
}

func NewMessage(u string, imgContent *image.Image) *Message {
	return &Message{userID: u, imgContent: imgContent}
}

func (bot *Bot) PostFile(cid string, filename string, content []byte) (*string, error) {
	// https://github.com/ras0q/traq-wordcloud-bot/blob/ac3445ca8d94588cf7407a08375a4ed91869c56e/pkg/traqapi/traqapi.go#L86
	// NOTE: go-traqがcontent-typeをapplication/octet-streamにしてしまうので自前でAPIを叩く
	// Ref: https://github.com/traPtitech/go-traq/blob/2c7a5f9aa48ef67a6bd6daf4018ca2dabbbbb2f3/client.go#L304
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)

	mh := make(textproto.MIMEHeader)
	mh.Set("Content-Type", "image/png")
	mh.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, filename))

	pw, err := mw.CreatePart(mh)
	if err != nil {
		return nil, fmt.Errorf("failed to create part: %w", err)
	}

	if _, err := pw.Write(content); err != nil {
		return nil, fmt.Errorf("failed to copy file: %w", err)
	}

	contentType := mw.FormDataContentType()
	mw.Close()

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("https://q.trap.jp/api/v3/files?channelId=%s", cid),
		&b,
	)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Authorization", "Bearer "+bot.auth.Value(traq.ContextAccessToken).(string))

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)

		return nil, fmt.Errorf("Error creating file: %s %s", res.Status, string(b))
	}

	var traqFile traq.FileInfo
	if err := json.NewDecoder(res.Body).Decode(&traqFile); err != nil {
		return nil, fmt.Errorf("Error decoding response: %w", err)
	}

	return &traqFile.Id, nil
}

func (bot *Bot) SendImage(cid string, msg *Message, embed bool) {
	buf := new(bytes.Buffer)

	if err := png.Encode(buf, *msg.imgContent); err != nil {
		log.Println("Failed To Encode png File")

		return
	}

	fileID, err := bot.PostFile(cid, "img.png", buf.Bytes())
	log.Println("FileID: " + *fileID)

	file := fmt.Sprintf("現在のユーザー: @%s\n\nhttps://q.trap.jp/files/%s", msg.userID, *fileID)
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
