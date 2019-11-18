package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	driver     = "postgres"
	dbhost     = "ec2-107-21-98-89.compute-1.amazonaws.com"
	dbuser     = "gydhbcepmvfojy"
	dbName     = "dccpn9636r8od"
	dbpassword = "4ce46cf1c6eabcd1d325dcd0fd31bddf3d573cdc6393f4e5077eead7fa3f53c8"
	sslmode    = "require"
)

type stmtConfig struct {
	stmt *sql.Stmt
	q    string
}

type Data struct {
	Db    *sql.DB
	Stmts map[string]*stmtConfig
}

var data Data

func InitDb() error {
	var err error
	data.Db, err = sql.Open(driver, fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s",
		dbhost, dbuser, dbName, dbpassword, sslmode))
	if err != nil {
		return err
	}
	data.Stmts = map[string]*stmtConfig{
		"listUser":   {q: "select * from \"User\";"},
		"getUser":    {q: "select * from \"User\" where id=?;"},
		"insertUser": {q: "Insert into \"User\" (id, name, deposit_addrs, receive_addrs, parent_id) values ($1,$2,$3,$4,$5);"},
		"updateUser": {q: "update User set name=? where id=?;"},
		"deleteUser": {q: "delete from User where id=?"},
		"listPlan":   {q: "select * from user_plan"},
		"getPlan":    {q: "select * from user_plan where user_id=?;"},
		"insertPlan": {q: "insert into \"user_plan\" (user_id, is_active, begin_date, invest) values ($1,$2,$3,$4);"},
		"updatePlan": {q: "update user_plan set is_active=false where (begin_date::date + '90 day'::interval)>?;"},
		"listTx":     {q: "select * from user_tx"},
		"getTx":      {q: "select * from user_tx where user_id=?;"},
		"insertTx":   {q: "insert into \"user_tx\" (user_id, is_deposit, amount, tx_id) values ($1,$2,$3,$4);"},
		// "updateTx":   {q: "update user_tx set ;"},
	}
	for k, v := range data.Stmts {
		data.Stmts[k].stmt, err = data.Db.Prepare(v.q)
	}
	return nil
}

func (d Data) Insert(u BotUser) (int64, error) {
	insertUser := data.Stmts["insertUser"].stmt

	_, err := insertUser.Exec(u.GetID(), u.GetName(), u.GetDepositAddress(),
		u.GetReceiveAddress(), u.GetParentId())
	if err != nil {
		return 0, err
	}
	for e := range u.GetActivePlans() {
		if err := data.insertPlan(u.GetID(), u.GetActivePlans()[e]); err != nil {
			return 0, err
		}
	}
	txs := append(u.GetDepositTransaction(), u.GetReceiveTransaction()...)
	for e := range txs {
		if err := data.insertTx(u.GetID(), txs[e]); err != nil {
			return 0, err
		}
	}

	return u.GetID(), nil
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

func (d Data) insertPlan(userId int64, plan UserPlan) error {
	insertPlan := d.Stmts["insertPlan"].stmt
	if _, err := insertPlan.Exec(userId, plan.IsActive(),
		plan.GetStartDate(), plan.GetAmount()); err != nil {
		return err
	}
	return nil
}

func (d Data) insertTx(userId int64, tx UserTransaction) error {
	insertTx := d.Stmts["insertTx"].stmt
	if _, err := insertTx.Exec(userId, tx.IsDepositTx(),
		tx.GetAmount(), tx.GetTxId()); err != nil {
		return err
	}
	return nil
}

func errUserNotFound(id int64) error {
	return fmt.Errorf("user with id: %d not found", id)
}
