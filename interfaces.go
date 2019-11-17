package main

import "time"

type Data interface {
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
	GetBalance() string
	GetParentId() int64
	SetParentId(parentId int64)
	GetReceiveTransaction() []UserTransaction
	GetDepositTransaction() []UserTransaction
	AddTransaction(transaction UserTransaction)
	GetActivePlans() []UserPlan
	AddPlan(plan UserPlan)
}

type UserTransaction interface {
	GetTxId() string
	SetTxId(txId string)
	IsDepositTx() bool
	Amount() string
}

type UserPlan interface {
	IsActive() bool
	GetStartDate() time.Time
	GetPlanType() PlanType
	SetPlanType(planType PlanType)
	GetAmount() float64
	SetAmount(investment float64)
}
