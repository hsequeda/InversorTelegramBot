package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"regexp"
)

const Exponent = 8

var (
	bot       *tgbotapi.BotAPI
	button    Button
	port      string
	key       string
	channelId string
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

	channelId = os.Getenv("CHNNL_ID")
	if channelId == "" {
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
		logrus.Error(err)
	}
	if info.LastErrorDate != 0 {
		logrus.Printf("[Telegram callback failed]%s", info.LastErrorMessage)
	}
	if err := InitDb(); err != nil {
		logrus.Error(err)
	}
}

func main() {
	updates := bot.ListenForWebhook("/InversorTelegramBot/")
	http.HandleFunc("/blockchain/", handleDeposit)
	go http.ListenAndServe("0.0.0.0:"+port, nil)

	for update := range updates {

		if update.Message == nil {
			if update.ChannelPost != nil {
				fmt.Println(update.ChannelPost)
			}
			continue
		}
		if update.Message.Chat.Type != "private" {
			continue
		}
		// user, err := GetUser(update.Message.Chat.ID)

		// Verify if message is "/start+parent_id message"
		if ok, _ := regexp.MatchString("(^/start [\\d#]+$)", update.Message.Text); ok {
			if !UserExist(update.Message.Chat.ID) {
				var userName = update.Message.From.UserName
				if userName == "" {
					userName = update.Message.From.FirstName
				}
				if err := AddUser(update.Message.Chat.ID,
					getParentIdFromMessage(update.Message.Text),
					userName); err != nil {
					logrus.Error(err)
				}
			}
		}
		user, err := GetUser(update.Message.Chat.ID)
		if err != nil {
			logrus.Error(err)
		}
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
			msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.UserName, text)
			if err != nil {
				logrus.Error(err)
			}
			_, err = bot.Send(msg)
			if err != nil {
				logrus.Fatal(err)
			}
			break
		case "Deposito":
			if user.GetDepositAddress() == "" {
				logrus.Info("Into deposit address")
				address, err := GetAddress()
				if err != nil {
					logrus.Error(err)
				}
				if err := SetAddrsToUser(update.Message.Chat.ID, address); err != nil {
					logrus.Error(err)
				}
				user, err = GetUser(user.GetID())
				if err != nil {
					logrus.Error(err)
				}
			}

			msg, err := button.InitButton(update.Message.Chat.ID, update.Message.From.UserName, fmt.Sprintf(
				"Envie la cantidad que desea invertir a la siguiente direccion: \n <code>%s</code> ",
				user.GetDepositAddress()))
			if err != nil {
				logrus.Error(err)
			}
			if _, err := bot.Send(msg); err != nil {
				logrus.Error(err)
			}

			break
		case user.GetName():
			break
		case "Movimientos", "Actualizar":
			text := "Aqui puede encontrar las ultimas 15 transacciones relacionadas con su cuenta."
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			if _, err := bot.Send(msg); err != nil {
				logrus.Fatal(err)
			}
			msg2, err := button.TransactionHistoryBtn(update.Message.Chat.ID)
			if err != nil {
				logrus.Error(err)
			}
			if _, err := bot.Send(msg2); err != nil {
				logrus.Fatal(err)
			}
			break
		case "Balance":
			user, err := GetUser(user.GetID())
			if err != nil {
				logrus.Fatal(err)
			}
			balance := decimal.New(user.GetBalance(), -Exponent)
			activeInversions, err := GetActiveInversions(user.GetID())

			if err != nil {
				logrus.Fatal(err)
			}
			text := fmt.Sprintf(
				"Saldo de la cuenta:\n"+
					"Saldo Extraible:\n"+
					"%s\n,"+
					"Inversiones Activas:\n"+
					"%s",
				balance.StringFixed(Exponent),
				activeInversions,
			)
			msg, err := button.InitButton(update.Message.Chat.ID, user.GetName(), text)
			if err != nil {
				logrus.Error(err)
			}
			_, err = bot.Send(msg)
			if err != nil {
				logrus.Fatal(err)
			}
			break
		case "Referidos":
			break
		default:
			inviteLink := fmt.Sprintf("https://t.me/Prebs_bot?start=%d", update.Message.Chat.ID)

			msg, err := button.InitButton(update.Message.Chat.ID, user.GetName(), inviteLink)
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
