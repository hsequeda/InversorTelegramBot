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
	return nil, nil
}

func AddInvestToUser(value string, userID int64) {
	// TODO
}
