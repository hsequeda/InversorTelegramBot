package main

type Data struct {
}

var data Data

func (d Data) Insert(user BotUser) (int64, error) {
	panic("implement me")
}

func (d Data) Get(id int64) (BotUser, error) {
	panic("implement me")
}

func (d Data) List() ([]BotUser, error) {
	panic("implement me")
}

func (d Data) Delete(id int64) error {
	panic("implement me")
}

func (d Data) Update(id int64, user BotUser) error {
	panic("implement me")
}
