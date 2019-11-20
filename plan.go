package main

import (
	"time"
)

type PlanType int

const (
	amountTypePlan1 = 11000 // In Satoshis
	amountTypePlan2 = 10000000
	amountTypePlan3 = 50000000
	MaxInvest       = 100000000
	Type1           = iota
	Type2
	Type3
)

type Plan struct {
	Start    time.Time
	Active   bool
	Type     PlanType
	Invested int64
	Id       int64
}

func (p *Plan) IsActive() bool {
	return p.Active
}

func (p *Plan) GetStartDate() time.Time {
	return p.Start
}

func (p *Plan) GetPlanType() PlanType {
	return p.Type
}

func (p *Plan) SetPlanType(planType PlanType) {
	p.Type = planType
}

func (p *Plan) GetAmount() int64 {
	return p.Invested
}

func (p *Plan) SetAmount(investment int64) {
}

func (p *Plan) GetId() int64 {
	return p.Id
}

func getPlanTypeForValue(value int64) PlanType {
	if value >= amountTypePlan1 && value < amountTypePlan2 {
		return amountTypePlan1
	}
	if value >= amountTypePlan2 && value < amountTypePlan3 {
		return amountTypePlan2
	}
	if value >= amountTypePlan3 {
		return amountTypePlan3
	}
	return -1
}
