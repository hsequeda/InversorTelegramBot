package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
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
	if v < 100000 {
		return fmt.Errorf("inversion amount is less than 0.0001 BTC")
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

// SetAddrsToUser replace deposit address of user.
func SetAddrsToUser(id int64, addr string) error {
	u, err := data.Get(id)
	if err != nil {
		return err
	}
	u.SetDepositAddress(addr)
	logrus.Info(u.GetDepositAddress())
	if err := data.Update(id, u); err != nil {
		return err
	}

	return nil
}

// UserExist verify if an user exist into database.
func UserExist(id int64) bool {
	_, err := data.Get(id)
	if err != sql.ErrNoRows {
		return true
	}
	return false
}

// AddTransactionToUser add a transaction to user.
func AddTransactionToUser(id int64, isDeposit bool, txId, value string) error {
	u, err := data.Get(id)
	if err != nil {
		return err
	}

	amount, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return err
	}

	u.AddTransaction(&Transaction{
		TxID:      txId,
		IsDeposit: isDeposit,
		Amount:    amount,
	})
	return nil
}

func AddUser(id, parentId int64, name string) error {
	logrus.Info("Add user")

	if id != parentId {
		if _, err := data.Insert(&User{Id: id, Name: name, ParentId: parentId}); err != nil {
			return err
		}
	} else {
		if _, err := data.Insert(&User{Id: id, ParentId: 0}); err != nil {
			return err
		}
	}
	return nil
}

func GetUser(id int64) (BotUser, error) {
	logrus.Info("Get user")
	return data.Get(id)
}
