package main

type PlanType int

func GetUserByAddress(address string) (*User, error) {
	// TODO
	return &User{
		Id:             1,
		Name:           "M. de Cervantes",
		DepositAddress: "",
		ReceiveAddress: "",
		Balance:        "",
		ParentId:       0,
	}, nil
}

func AddInvestToUser(value string, userID int64) {
	// TODO
}
