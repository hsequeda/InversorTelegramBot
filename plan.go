package main

import (
	"time"
)

type Plan struct {
}

func (p Plan) IsActive() bool {
	panic("implement me")
}

func (p Plan) GetStartDate() time.Time {
	panic("implement me")
}

func (p Plan) GetPlanType() PlanType {
	panic("implement me")
}

func (p Plan) SetPlanType(planType PlanType) {
	panic("implement me")
}

func (p Plan) GetAmount() float64 {
	panic("implement me")
}

func (p Plan) SetAmount(investment float64) {
	panic("implement me")
}
