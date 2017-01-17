package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
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
		replyToken := event.ReplyToken
		switch event.Type {
			case linebot.EventTypeMessage:
				switch message := event.Message.(type) {
					case *linebot.TextMessage:
						profile, getProfileErr := bot.GetProfile(event.Source.UserID).Do()
				if getProfileErr != nil {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(getProfileErr.Error()))
					log.Println(getProfileErr)
				}

				cmd := strings.Fields(message.Text)

				switch cmd[0] {
					case "加入":
					
						text := profile.DisplayName + " 您好，已將您加入傳送對象，未來將會傳送天氣警報資訊給您 ^＿^ "
						if _, replyErr := bot.ReplyMessage(
							replyToken,
							linebot.NewTextMessage(text)).Do(); replyErr != nil {
							log.Println(replyErr)
						}
					}
				}
			}				
		}
	}

