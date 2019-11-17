package main

type Transaction struct {
	TxID      string
	IsDeposit bool
}

func (t *Transaction) GetTxId() string {
	return t.TxID
}

func (t *Transaction) SetTxId(txId string) {
	t.TxID = txId
}

func (t *Transaction) IsDepositTx() bool {
	panic("implement me")
}

func (t *Transaction) Amount() string {
	panic("implement me")
}
