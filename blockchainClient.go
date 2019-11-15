package main

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"net/http"
)

// type Blockchain struct{}

type TicketProperty struct {
	Fifteen float32 `json:"15m"`
	Last    float32 `json:"last"`
	Buy     float32 `json:"buy"`
	Sell    float32 `json:"sell"`
	Symbol  string  `json:"symbol"`
}

func GetPrices() (map[string]TicketProperty, error) {

	resp, err := http.Get("https://blockchain.info/ticker")
	if err != nil {
		return nil, err
	}
	var tickets = make(map[string]TicketProperty)
	if err := json.NewDecoder(resp.Body).Decode(&tickets); err != nil {
		return nil, err
	}
	return tickets, nil
}

func GetAddress() (string, error) {
	req, err := http.NewRequest("POST", "https://www.blockonomics.co/api/new_address", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", key))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	var address = make(map[string]string)
	err = json.NewDecoder(resp.Body).Decode(&address)
	return address["address"], err
}

func handleDeposit(w http.ResponseWriter, r *http.Request) {
	fmt.Println(channel_id)
	status := r.URL.Query().Get("status")
	address := r.URL.Query().Get("addr")
	value := r.URL.Query().Get("value")
	txid := r.URL.Query().Get("txid")
	switch status {
	case "0":
		u, err := GetUserByAddress(address)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(
			fmt.Sprintf("Se ha detectado una tranccion "+
				"no confirmada por parte del usuario %s.\n"+
				"Transaction ID:%s", u.Name, txid))
		break
	case "1":
		u, err := GetUserByAddress(address)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Info(
			fmt.Sprintf("Se ha detectado una tranccion "+
				"parcialmente confirmada por parte del usuario %s.\n"+
				"Transaction ID:%s", u.Name, txid))
		break
	case "2":
		u, err := GetUserByAddress(address)
		if err != nil {
			logrus.Error(err)
		}
		AddInvestToUser(value, u.Id)

		msg := tgbotapi.NewMessage(channel_id,
			fmt.Sprintf("Nueva inversion:\n "+
				"%s ha invertido %s BTC!\n"+
				"Transaction ID:\n"+
				"%s", u.Name, value, txid))
		if _, err := bot.Send(msg); err != nil {
			logrus.Error(err)
		}
		break
	}
}
