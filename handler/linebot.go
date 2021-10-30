package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type lineBotHandler struct {
	bot *linebot.Client
}

func NewLineBotHandler(bot *linebot.Client) lineBotHandler {
	return lineBotHandler{bot: bot}
}

func (line lineBotHandler) HandleWebHook(w http.ResponseWriter, r *http.Request) {
	events, err := line.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			replyToken := event.ReplyToken
			userID := event.Source.UserID
			groupID := event.Source.GroupID
			RoomID := event.Source.RoomID
			fmt.Printf("replyToken:%s\nuserID:%s\ngroupID:%s\nRoomID:%s", replyToken, userID, groupID, RoomID)
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				handleTextMessage(line, event, message)
			case *linebot.StickerMessage:
				handleStickerMessage(line, event, message)
			case *linebot.ImageMessage:
				handleImageMessage(line, event, message)
			}
		}
	}
}

func handleTextMessage(line lineBotHandler, event *linebot.Event, message *linebot.TextMessage) {
	if _, err := line.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
		log.Print(err)
	}
}

func handleStickerMessage(line lineBotHandler, event *linebot.Event, message *linebot.StickerMessage) {
	fmt.Println(message.Keywords)
	replyMessage := fmt.Sprintf(
		"Image id is %s", message.ID)
	if _, err := line.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
	}
}

func handleImageMessage(line lineBotHandler, event *linebot.Event, message *linebot.ImageMessage) {
	replyMessage := fmt.Sprintf(
		"sticOriginalContentURLker  is %s, OriginalContentURL is %s", message.PreviewImageURL, message.PreviewImageURL)
	if _, err := line.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
	}
}
