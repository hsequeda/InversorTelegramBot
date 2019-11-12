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

	// _, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://inversortelegrambot.herokuapp.com/",""))
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
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		b := tgbotapi.NewKeyboardButton("Button")
		bc := tgbotapi.NewKeyboardButtonContact("Button contact")
		bl := tgbotapi.NewKeyboardButtonLocation("Button location")
		br := tgbotapi.NewKeyboardButtonRow(b, bc, bl)
		rk := tgbotapi.NewReplyKeyboard(br)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "d ")
		msg.ReplyMarkup = rk
		rm, err := bot.Send(msg)
		if err != nil {
			fmt.Println(err)
		}
		log.Printf("%+v\n", rm)
	}
}
