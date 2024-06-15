package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/traP-jp/h24s_17-backend/bot"
	traqbot "github.com/traPtitech/traq-bot"
)

func (s *State) SetupRoutes(e *echo.Echo, b *bot.Bot) {
	e.GET("/hello/:name", s.HelloHandler)

	e.POST("", MakeBotHandler(b.VerificationToken, b.MakeHandlers()))
	api := e.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}

func MakeBotHandler(vt string, handlers traqbot.EventHandlers) func(c echo.Context) error {
	// https://github.com/traPtitech/traq-bot/blob/a4242da6299790b6f5d1cac6afd7ea351a7ce852/server.go#L45-L114
	return func(c echo.Context) error {
		req := c.Request()
		// VerificationTokenチェック
		if req.Header.Get("X-TRAQ-BOT-TOKEN") != vt {
			return c.NoContent(http.StatusForbidden)
		}
		// Eventヘッダチェック
		event := req.Header.Get("X-TRAQ-BOT-EVENT")
		if len(event) == 0 {
			return c.NoContent(http.StatusBadRequest)
		}
		// RequestがJSONかどうか
		if !strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
			return c.NoContent(http.StatusBadRequest)
		}
		var payload interface{}
		switch event {
		case traqbot.Ping:
			payload = &traqbot.PingPayload{}
		case traqbot.Joined:
			payload = &traqbot.JoinedPayload{}
		case traqbot.Left:
			payload = &traqbot.LeftPayload{}
		case traqbot.MessageCreated:
			payload = &traqbot.MessageCreatedPayload{}
		case traqbot.MessageUpdated:
			payload = &traqbot.MessageUpdatedPayload{}
		case traqbot.MessageDeleted:
			payload = &traqbot.MessageDeletedPayload{}
		case traqbot.BotMessageStampsUpdated:
			payload = &traqbot.BotMessageStampsUpdatedPayload{}
		case traqbot.DirectMessageCreated:
			payload = &traqbot.DirectMessageCreatedPayload{}
		case traqbot.DirectMessageUpdated:
			payload = &traqbot.DirectMessageUpdatedPayload{}
		case traqbot.DirectMessageDeleted:
			payload = &traqbot.DirectMessageDeletedPayload{}
		case traqbot.ChannelCreated:
			payload = &traqbot.ChannelCreatedPayload{}
		case traqbot.ChannelTopicChanged:
			payload = &traqbot.ChannelTopicChangedPayload{}
		case traqbot.UserCreated:
			payload = &traqbot.UserCreatedPayload{}
		case traqbot.StampCreated:
			payload = &traqbot.StampCreatedPayload{}
		case traqbot.TagAdded:
			payload = &traqbot.TagAddedPayload{}
		case traqbot.TagRemoved:
			payload = &traqbot.TagRemovedPayload{}
		default:
			return c.NoContent(http.StatusNotImplemented)
		}
		if err := json.NewDecoder(req.Body).Decode(payload); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		h, ok := handlers[event]
		if ok && h != nil {
			h(event, payload)
		}
		return c.NoContent(http.StatusNoContent)
	}
}
