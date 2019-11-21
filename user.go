package main

type User struct {
	Id             int64
	Name           string
	DepositAddress string
	ReceiveAddress string
	Balance        int64
	ParentId       int64
	Txs            []UserTransaction
	Plans          []UserPlan
}

func (u *User) SetPlans(plans []UserPlan) {
	u.Plans = plans
}

func (u *User) GetPlans() []UserPlan {
	return u.Plans
}

func (u *User) GetID() int64 {
	return u.Id
}

func (u *User) SetID(id int64) {
	u.Id = id
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetDepositAddress() string {
	return u.DepositAddress
}

func (u *User) SetDepositAddress(addr string) {
	u.DepositAddress = addr
}

func (u *User) GetReceiveAddress() string {
	return u.ReceiveAddress
}

func (u *User) SetReceiveAddress(addr string) {
	u.ReceiveAddress = addr
}

func (u *User) GetBalance() int64 {
	return u.Balance
}

func (u *User) SetBalance(newBalance int64) {
	u.Balance = newBalance
}

func (u *User) GetParentId() int64 {
	return u.ParentId
}

func (u *User) SetParentId(parentId int64) {
	u.ParentId = parentId
}

func (u *User) GetReceiveTransaction() []UserTransaction {
	var txs []UserTransaction
	for e := range u.Txs {
		if !u.Txs[e].IsDepositTx() {
			txs = append(txs, u.Txs[e])
		}
	}
	return txs
}

func (u *User) GetDepositTransaction() []UserTransaction {
	var txs []UserTransaction
	for e := range u.Txs {
		if u.Txs[e].IsDepositTx() {
			txs = append(txs, u.Txs[e])
		}
	}
	return txs
}

func (u *User) AddTransaction(transaction UserTransaction) {
	u.Txs = append(u.Txs, transaction)
}

func (u *User) GetActivePlans() []UserPlan {
	var plans []UserPlan
	for e := range u.Plans {
		if u.Plans[e].GetEndDate().Before(u.Plans[e].GetLastPaymentDate()) {
			plans = append(plans, u.Plans[e])
		}
	}
	return plans
}

func (u *User) AddPlan(plan UserPlan) {
	u.Plans = append(u.Plans, plan)
}
