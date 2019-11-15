package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

var (
	bot        *tgbotapi.BotAPI
	button     Button
	port       string
	key        string
	channel_id int64
)

func init() {
	var err error
	key = os.Getenv("APIKEY")
	if key == "" {
		logrus.Error("$APIKEY is empty")
	}
	port = os.Getenv("PORT")
	if key == "" {
		logrus.Error("$PORT is empty")
	}

	channel_id, err = strconv.ParseInt(os.Getenv("CHNNL_ID"), 10, 0)
	if err != nil || channel_id == 0 {
		logrus.Error("error obtains $CHNNL_ID, can by empty")
	}
	if key == "" {
		logrus.Error("$PORT is empty")
	}
	bot, err = tgbotapi.NewBotAPI("931110470:AAHmRc3jqseVa8W5qTrgjueR6HhU0PIOuTI")

	if err != nil {
		logrus.Error(err)
	}
	bot.Debug = true
	info, err := bot.GetWebhookInfo()
	if err != nil {
		logrus.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		logrus.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
}

func main() {
	updates := bot.ListenForWebhook("/InversorTelegramBot/")
	http.HandleFunc("/blockchain/", showData)
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		switch update.Message.Text {
		case "Precios":
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
			break
		case "Deposito":
			if !UserExist(update.Message.Chat.ID) {
				address, err := GetAddress()
				SetAddrsToUser(address)
				if err != nil {
					logrus.Error(err)
				}
				msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.FirstName, fmt.Sprintf(
					" Envie la cantidad que desea invertir a la siguiente direccion: \n %s", address))
				if err != nil {
					logrus.Error(err)
				}
				if _, err := bot.Send(msg); err != nil {
					logrus.Error(err)
				}
			}
			break
		default:
			inviteLink, err := bot.GetInviteLink(tgbotapi.ChatConfig{
				ChatID:             931110470,
				SuperGroupUsername: "aa",
			})
			if err != nil {
				logrus.Error(err)
			}
			msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.FirstName, inviteLink)
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

func SetAddrsToUser(s string) {
	// TODO

}

func UserExist(i int64) bool {
	// TODO
	return false
}
