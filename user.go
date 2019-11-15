package main

type User struct {
	Id             int64
	Name           string
	DepositAddress string
	ReciveAddress  string
	Balance        string
	ParentId       int64
}

func GetUserByAddress(address string) (*User, error) {
	// TODO
	return &User{
		Id:             1,
		Name:           "M. de Cervantes",
		DepositAddress: "",
		ReciveAddress:  "",
		Balance:        "",
		ParentId:       0,
	}, nil
}

func AddInvestToUser(value string, userID int64) {
	// TODO
}
