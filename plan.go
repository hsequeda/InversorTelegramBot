package main

import (
	"time"
)

type PlanType int

const (
	amountTypePlan1 = 100000 // In Satoshis
	amountTypePlan2 = 10000000
	amountTypePlan3 = 50000000
	// MaxInvest       = 100000000
	Type1 = iota
	Type2
	Type3
)

type Plan struct {
	Start       time.Time
	End         time.Time
	LastPayment time.Time
	Invested    int64
	Id          int64
}

func (p *Plan) GetLastPaymentDate() time.Time {
	return p.LastPayment
}

func (p *Plan) SetLastPaymentDate(date time.Time) {
	p.LastPayment = date
}

func (p *Plan) GetStartDate() time.Time {
	return p.Start
}

func (p *Plan) GetPlanType() PlanType {
	return getPlanTypeForValue(p.GetAmount())
}

func (p *Plan) GetAmount() int64 {
	return p.Invested
}

func (p *Plan) GetId() int64 {
	return p.Id
}

func (p *Plan) GetEndDate() time.Time {
	return p.End
}

func getPlanTypeForValue(value int64) PlanType {
	if value >= amountTypePlan1 && value < amountTypePlan2 {
		return Type1
	}
	if value >= amountTypePlan2 && value < amountTypePlan3 {
		return Type2
	}
	if value >= amountTypePlan3 {
		return Type3
	}
	return -1
}
