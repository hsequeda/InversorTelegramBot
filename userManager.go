package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetUserByDepositAddress
func GetUserByDepositAddress(address string) (BotUser, error) {
	userList, err := data.List()
	if err != nil {
		return nil, err
	}
	for e := range userList {
		if userList[e].GetDepositAddress() == address {
			return userList[e], err
		}
	}
	return nil, errors.New(fmt.Sprintf("user with address %s not found", address))
}

// AddInvestToUser add a invest plan to user by userID.
func AddInvestToUser(value string, userID int64) error {
	user, err := data.Get(userID)
	if err != nil {
		return err
	}
	v, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return err
	}
	p := createPlan(v)
	if p != nil {
		user.AddPlan(p)
		if err := data.Update(user.GetID(), user); err != nil {
			return err
		}
	}
	return nil
}

func createPlan(value int64) *Plan {
	planType := getPlanTypeForValue(value)
	if planType != -1 {
		return &Plan{
			Start:    time.Now(),
			Active:   true,
			Type:     planType,
			Invested: value,
		}
	}
	return nil
}

func SetAddrsToUser(s string) {
	// TODO

}

// UserExist verify if an user exist into database.
func UserExist(id int64) bool {
	_, err := data.Get(id)
	if err != errUserNotFound(id) {
		return true
	}
	return false
}
