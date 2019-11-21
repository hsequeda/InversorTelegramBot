package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"time"
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
		"getUser":    {q: "select * from \"User\" where id=$1;"},
		"insertUser": {q: "Insert into \"User\" (id, name, deposit_addrs, receive_addrs, parent_id) values ($1,$2,$3,$4,$5);"},
		"updateUser": {q: "update \"User\" set name=$1,deposit_addrs=$2, receive_addrs=$3 where id=$4;"},
		"deleteUser": {q: "delete from \"User\" where id=$1"},
		"listPlan":   {q: "select * from \"user_plan\""},
		"getPlan":    {q: "select * from \"user_plan\" where user_id=$1;"},
		"insertPlan": {q: "insert into \"user_plan\" (user_id, is_active, begin_date, invest) values ($1,$2,$3,$4);"},
		// "updatePlan": {q: "update \"user_plan\" set is_active=false where (begin_date::date + '90 day'::interval)>?;"},
		// "updatePlan": {q: "update \"user_plan\" set is_active=false where plan_id=$1;"},
		"listTx":   {q: "select * from user_tx"},
		"getTx":    {q: "select * from \"user_tx\" where user_id=$1;"},
		"insertTx": {q: "insert into \"user_tx\" (user_id, is_deposit, amount, tx_id) values ($1,$2,$3,$4);"},
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
	getUser := d.Stmts["getUser"].stmt
	u := User{}
	if err := getUser.QueryRow(id).
		Scan(&u.Id, &u.Name, &u.DepositAddress, &u.ReceiveAddress, &u.ParentId); err != nil {
		return nil, err
	}

	// if err := d.updateDatePlans(); err != nil {
	// 	return nil, err
	// }

	p, err := data.getPlans(u.Id)
	if err != nil {
		return nil, err
	}
	u.Plans = append(u.Plans, p...)

	txs, err := data.getTxs(u.Id)
	if err != nil {
		return nil, err
	}
	u.Txs = append(u.Txs, txs...)

	return &u, nil
}

func (d Data) List() ([]BotUser, error) {
	listUser := d.Stmts["listUser"].stmt
	rows, err := listUser.Query()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users = make([]BotUser, 0)
	for rows.Next() {
		var u = User{}
		if err := rows.Scan(&u.Id, &u.Name, &u.DepositAddress, &u.ReceiveAddress, &u.ParentId); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (d Data) Delete(id int64) error {
	delUser := d.Stmts["deleteUser"].stmt
	_, err := delUser.Exec(id)
	return err
}

func (d Data) Update(id int64, user BotUser) error {
	logrus.Info("Update")
	updUser := d.Stmts["updateUser"].stmt
	_, err := updUser.Exec(user.GetName(), user.GetDepositAddress(), user.GetReceiveAddress(), user.GetID())
	if err != nil {
		return err
	}
	plans, err := d.getPlans(id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	for _, uPlan := range user.GetActivePlans() {
		exist := false
		for _, plan := range plans {
			if plan.IsActive() && plan.GetId() == uPlan.GetId() {
				exist = true
				if err := d.updatePlan(uPlan); err != nil {
					return err
				}
			}
			if !exist {
				if err := d.insertPlan(user.GetID(), uPlan); err != nil {
					return err
				}
				exist = false
			}
		}
	}

	txs, err := d.getTxs(id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	userTxs := append(user.GetReceiveTransaction(), user.GetDepositTransaction()...)
	for _, uTx := range userTxs {
		exist := false
		for _, tx := range txs {
			if uTx.GetTxId() == tx.GetTxId() {
				exist = true
				if err := d.updateTx(uTx); err != nil {
					return err
				}
			}
			if !exist {
				if err := d.insertTx(user.GetID(), uTx); err != nil {
					return err
				}
				exist = false
			}
		}
	}

	return nil
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

func (d Data) getPlans(userId int64) ([]UserPlan, error) {
	getPlans := d.Stmts["getPlan"].stmt
	rows, err := getPlans.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plans := make([]UserPlan, 0)
	for rows.Next() {
		p := Plan{}

		if err := rows.Scan(&userId, &p.Active, &p.Start, &p.Invested, &p.Id); err != nil {
			return nil, err
		}
		plans = append(plans, &p)
	}
	return plans, nil
}

func (d Data) getTxs(userId int64) ([]UserTransaction, error) {
	getTxs := data.Stmts["getTx"].stmt

	rows, err := getTxs.Query(userId)
	if err != nil {
		return nil, err
	}
	txs := make([]UserTransaction, 0)
	defer rows.Close()

	for rows.Next() {
		tx := Transaction{}
		if err := rows.Scan(&userId, &tx.IsDeposit, &tx.Amount, &tx.TxID); err != nil {
			return nil, err
		}
		txs = append(txs, &tx)
	}
	return txs, nil
}

func (d *Data) updateDatePlans() error {
	updateDatePlans := data.Stmts["updatePlan"].stmt
	if _, err := updateDatePlans.Exec(time.Now()); err != nil {
		return err
	}
	return nil
}

func (d Data) updatePlan(plan UserPlan) error {
	// 	updtPlanStmt := d.Stmts["updatePlan"].stmt
	// updtPlanStmt.Exec()
	// TODO
	return nil
}

func (d Data) updateTx(uTx UserTransaction) error {

	return nil
}

func errUserNotFound(id int64) error {
	return fmt.Errorf("user with id: %d not found", id)
}
