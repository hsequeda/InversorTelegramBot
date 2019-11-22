package main

import "time"

type DbManager interface {
	Insert(user BotUser) (int64, error)
	Get(id int64) (BotUser, error)
	List() ([]BotUser, error)
	Delete(id int64) error
	Update(id int64, user BotUser) error
}

type BotUser interface {
	GetID() int64
	SetID(id int64)
	GetName() string
	SetName(name string)
	GetDepositAddress() string
	SetDepositAddress(addr string)
	GetReceiveAddress() string
	SetReceiveAddress(addr string)
	GetBalance() int64
	SetBalance(int64)
	GetRefersBonus() int64
	SetRefersBonus(value int64)
	GetParentId() int64
	SetParentId(parentId int64)
	GetReceiveTransaction() []UserTransaction
	GetDepositTransaction() []UserTransaction
	AddTransaction(transaction UserTransaction)
	GetActivePlans() []UserPlan
	GetPlans() []UserPlan
	SetPlans(plan []UserPlan)
}

type UserTransaction interface {
	GetTxId() string
	SetTxId(txId string)
	IsDepositTx() bool
	GetAmount() int64
	GetDate() time.Time
}

type UserPlan interface {
	GetId() int64
	GetStartDate() time.Time
	GetLastPaymentDate() time.Time
	SetLastPaymentDate(date time.Time)
	GetEndDate() time.Time
	GetPlanType() PlanType
	GetAmount() int64
}
