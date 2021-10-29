package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {

	// Connect to Line Bot
	ConnectLinebot()

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/webhook", handleWebHook)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "4325"
		log.Printf("defaulting to port %s", port)
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// Connect to Line Bot
func ConnectLinebot() {
	log.Print("🚀 starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebHook(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
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
				handleTextMessage(event, message)
			case *linebot.StickerMessage:
				handleStickerMessage(event, message)
			case *linebot.ImageMessage:
				handleImageMessage(event, message)
			}

		}
	}
}

func handleTextMessage(event *linebot.Event, message *linebot.TextMessage) {
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
		log.Print(err)
	}
}

func handleStickerMessage(event *linebot.Event, message *linebot.StickerMessage) {
	fmt.Println(message.Keywords)
	replyMessage := fmt.Sprintf(
		"Image id is %s", message.ID)
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
	}
}

func handleImageMessage(event *linebot.Event, message *linebot.ImageMessage) {
	replyMessage := fmt.Sprintf(
		"sticOriginalContentURLker  is %s, OriginalContentURL is %s", message.PreviewImageURL, message.PreviewImageURL)
	if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
		log.Print(err)
	}
}
