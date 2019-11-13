package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Button struct{}

func (*Button) InitButton(id int64, userName, msg string) (tgbotapi.Chattable, error) {
	var initReplyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Balance"), tgbotapi.NewKeyboardButton("Movimientos"),
			tgbotapi.NewKeyboardButton(userName), tgbotapi.NewKeyboardButton("Deposito"),
			tgbotapi.NewKeyboardButton("Retirar"), tgbotapi.NewKeyboardButton("Referidos"),
			tgbotapi.NewKeyboardButton("Extra"), tgbotapi.NewKeyboardButton(fmt.Sprintf("Prices"))))
	sendMsg := tgbotapi.NewMessage(id, msg)
	sendMsg.ReplyMarkup = initReplyKeyboard
	return sendMsg, nil
}
