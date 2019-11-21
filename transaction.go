package main

import "time"

type Transaction struct {
	TxID      string
	IsDeposit bool
	Amount    int64
	Date      time.Time
}

func (t *Transaction) GetDate() time.Time {
	return t.Date
}

func (t *Transaction) GetTxId() string {
	return t.TxID
}

func (t *Transaction) SetTxId(txId string) {
	t.TxID = txId
}

func (t *Transaction) IsDepositTx() bool {
	return t.IsDeposit
}

func (t *Transaction) GetAmount() int64 {
	return t.Amount
}
