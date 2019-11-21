package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"sort"
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
	sendMsg.ParseMode = "html"
	return sendMsg, nil
}

func (button2 *Button) TransactionHistoryBtn(id int64) (tgbotapi.Chattable, error) {
	var transacitonHistoryKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Update"), tgbotapi.NewKeyboardButton("Atras")))

	user, err := GetUser(id)
	if err != nil {
		return nil, err
	}
	transactions := append(user.GetDepositTransaction(), user.GetReceiveTransaction()...)
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].GetDate().After(transactions[j].GetDate())
	})
	var msg = "<b>Transaction ID:</b>\n"
	for e := range transactions {
		if e > 15 {
			break
		}
		msg += fmt.Sprintf("<a href=\"https://blockchain.info/tx/%s\">%s</a>\n",
			transactions[e].GetTxId(), transactions[e].GetTxId())
	}
	sendMsg := tgbotapi.NewMessage(id, msg)
	sendMsg.ReplyMarkup = transacitonHistoryKeyboard
	sendMsg.ParseMode = "html"
	return sendMsg, nil
}
