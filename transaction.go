package main

type Transaction struct {
}

func (t Transaction) GetTxId() string {
	panic("implement me")
}

func (t Transaction) SetTxId(txId string) {
	panic("implement me")
}

func (t Transaction) IsDepositTx() bool {
	panic("implement me")
}

func (t Transaction) Amount() string {
	panic("implement me")
}
