package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	bot, err := tgbotapi.NewBotAPI("931110470:AAHmRc3jqseVa8W5qTrgjueR6HhU0PIOuTI")
	if err != nil {
		fmt.Println(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://inversortelegrambot.herokuapp.com/", "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServeTLS("0.0.0.0:"+port, "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
