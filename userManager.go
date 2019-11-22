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
	logrus.Info("Add Invest")
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
		user.SetPlans(append(user.GetActivePlans(), p))
		if err := data.Update(user.GetID(), user); err != nil {
			return err
		}
		if user.GetParentId() != 0 {
			if err := SetBonusToParent(1, user.GetParentId(), p.GetAmount()); err != nil {
				return err
			}
		}

	}
	return nil
}

func SetBonusToParent(lv int, id int64, amount int64) error {
	user, err := data.Get(id)
	if err != nil {
		return err
	}
	switch lv {
	case 1:
		user.SetRefersBonus(int64(float64(user.GetRefersBonus()) + float64(amount)*0.03))
		user.SetBalance(user.GetBalance() + user.GetRefersBonus())
		break
	case 2:
		user.SetRefersBonus(int64(float64(user.GetRefersBonus()) + float64(amount)*0.02))
		user.SetBalance(user.GetBalance() + user.GetRefersBonus())
		break
	case 3:
		user.SetRefersBonus(int64(float64(user.GetRefersBonus()) + float64(amount)*0.01))
		user.SetBalance(user.GetBalance() + user.GetRefersBonus())
		break
	}
	if err := data.Update(user.GetID(), user); err != nil {
		return err
	}
	if user.GetParentId() != 0 && lv != 3 {
		if err := SetBonusToParent(lv+1, user.GetParentId(), amount); err != nil {
			return err
		}
	}
	return nil
}

func createPlan(value int64) *Plan {
	planType := getPlanTypeForValue(value)
	if planType != -1 {
		return &Plan{
			Start:       getDate(),
			End:         getDate().Add(90 * time.Hour * 24),
			LastPayment: getDate(),
			Invested:    value,
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
		Date:      getDate(),
	})

	if err := data.Update(u.GetID(), u); err != nil {
		return err
	}
	return nil
}

func AddUser(id, parentId int64, name string) error {
	logrus.Info("Add user")

	if id != parentId {
		if _, err := data.Insert(&User{Id: id, Name: name, ParentId: parentId}); err != nil {
			return err
		}
	} else {
		if _, err := data.Insert(&User{Id: id, Name: name, ParentId: 0}); err != nil {
			return err
		}
	}
	return nil
}

func GetUser(id int64) (BotUser, error) {
	logrus.Info("Get user")
	user, err := data.Get(id)
	if err != nil {
		return nil, err
	}

	if len(user.GetActivePlans()) > 0 {
		var updtPlans = make([]UserPlan, 0)
		for _, v := range user.GetActivePlans() {
			var durationToPay time.Duration
			if getDate().Equal(v.GetEndDate()) || getDate().After(v.GetEndDate()) {
				durationToPay = v.GetEndDate().Sub(v.GetLastPaymentDate())
			} else {
				durationToPay = getDate().Sub(v.GetLastPaymentDate())
			}
			dayToPay := int64(durationToPay.Hours()) / 24
			switch v.GetPlanType() {
			case Type1:
				user.SetBalance(user.GetBalance() + int64(float32(v.GetAmount()*dayToPay)*0.03))
				break
			case Type2:
				user.SetBalance(user.GetBalance() + int64(float32(v.GetAmount()*dayToPay)*0.035))
				break
			case Type3:
				user.SetBalance(user.GetBalance() + int64(float64(v.GetAmount()*dayToPay)*0.038))
				break
			}
			v.SetLastPaymentDate(getDate())
			updtPlans = append(updtPlans, v)
		}
		user.SetPlans(updtPlans)
		if err := data.Update(user.GetID(), user); err != nil {
			return nil, err
		}
	}
	return data.Get(id)
}

func GetActiveInversions(id int64) (int64, error) {
	user, err := data.Get(id)
	if err != nil {
		return 0, err
	}
	plans := user.GetActivePlans()
	var inversions int64
	for e := range plans {
		inversions += plans[e].GetAmount()
	}

	return inversions, nil
}

func GetTotalProfit(id int64) (int64, error) {
	user, err := data.Get(id)
	if err != nil {
		return 0, err
	}
	plans := user.GetActivePlans()
	var inversions int64
	for e := range plans {
		switch plans[e].GetPlanType() {
		case Type1:
			inversions += int64(float64(plans[e].GetAmount()) * 0.03 * 90)
			break
		case Type2:
			inversions += int64(float64(plans[e].GetAmount()) * 0.035 * 90)
			break
		case Type3:
			inversions += int64(float64(plans[e].GetAmount()) * 0.038 * 90)
			break
		}
	}
	return inversions, nil
}
