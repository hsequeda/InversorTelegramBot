package main

import (
	"fmt"
	_ "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":"+port, http.DefaultServeMux)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

	http.HandleFunc("/", helloworld)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Hello, World!")); err != nil {
		fmt.Println(err)
	}
}
