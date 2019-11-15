package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Button struct{}

func (*Button) InitButton(id int64, userName, msg string) (tgbotapi.Chattable, error) {
	var initReplyKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Balance")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Movimientos"), tgbotapi.NewKeyboardButton(userName)),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Deposito"), tgbotapi.NewKeyboardButton("Retirar")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Referidos"), tgbotapi.NewKeyboardButton("Extra")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Precios")))

	sendMsg := tgbotapi.NewMessage(id, msg)
	sendMsg.ReplyMarkup = initReplyKeyboard
	return sendMsg, nil
}
