package main

import (
	"encoding/json"
	"fmt"
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
