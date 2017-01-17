package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
//	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/callback", callbackHandlerExample)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
	
}

func callbackHandlerExample(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)


  if err != nil {
 
  }
  if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("hello")).Do(); 
	err != nil {
 
  }
	
}



/*
//回復程序者
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
	//	if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" 成功接受到訊息!!")).Do(); 
	//	err != nil {
	//		log.Print(err)
	//	}
//-------------回復訊息 example--------------
//	bot, err := linebot.New(<channel secret>, <channel token>)
//  if err != nil {
//  ...
//  }
//  if _, err := bot.ReplyMessage(<replyToken>, linebot.NewTextMessage("hello")).Do(); err != nil {
// ...
//  }
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
			
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); 
				err != nil {
					log.Print(err)
				}
				
			}
		}
	}
	*/
}
