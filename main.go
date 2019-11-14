package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var button Button

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
		logrus.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		logrus.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	http.HandleFunc("/blockchain/", showData)
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		logrus.Infof("%#v", update)
		if update.Message.Text == "Prices" {
			prices, err := GetPrices()
			if err != nil {
				logrus.Error(err)
			}
			text := ""
			for k, v := range prices {
				text += fmt.Sprintf("BTC to %s: %s%f\n", k, v.Symbol, v.Last)
			}
			msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.FirstName, text)
			if err != nil {
				logrus.Error(err)
			}
			_, err = bot.Send(msg)
			if err != nil {
				logrus.Fatal(err)
			}
		} else {
			msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.FirstName, "Welcome")
			if err != nil {
				logrus.Error(err)
			}
			_, err = bot.Send(msg)
			if err != nil {
				logrus.Fatal(err)
			}
		}
	}
}

func showData(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Print(err)
	}
	logrus.Print(string(b))
}
