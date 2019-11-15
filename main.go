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

var (
	button Button
	port   string
	key    string
)

func init() {
	key = os.Getenv("APIKEY")
	if key == "" {
		logrus.Error("$APIKEY is empty")
	}
	port = os.Getenv("PORT")
	if key == "" {
		logrus.Error("$PORT is empty")
	}

}

func main() {
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
	updates := bot.ListenForWebhook("/InversorTelegramBot/")
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {
		switch update.Message.Text {
		case "Price":
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

func SetAddrsToUser(s string) {
	// TODO
}

func UserExist(i int64) bool {
	// TODO
	return false
}

func showData(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Print(err)
	}
	status := r.URL.Query().Get("status")
	logrus.Info("status: ", status)
	address := r.URL.Query().Get("addr")
	logrus.Info("address: ", address)
	value := r.URL.Query().Get("value")
	logrus.Info("value: ", value)

	txid := r.URL.Query().Get("txid")
	logrus.Info("txid: ", txid)

	logrus.Print(string(b))
}
