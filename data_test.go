package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// func TestData_Delete(t *testing.T) {
// 	type fields struct {
// 		Db    *sql.DB
// 		Stmts map[string]*stmtConfig
// 	}
// 	type args struct {
// 		id int64
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := Data{
// 				Db:    tt.fields.Db,
// 				Stmts: tt.fields.Stmts,
// 			}
// 			if err := d.Delete(tt.args.id); (err != nil) != tt.wantErr {
// 				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestData_Get(t *testing.T) {
	type fields struct {
		u BotUser
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    BotUser
		wantErr bool
	}{
		{name: "ok",
			fields: fields{&User{
				Id:             1,
				Name:           "Walter",
				DepositAddress: "ssssss",
				ReceiveAddress: "xxxxxxxx",
				Balance:        "12345",
				ParentId:       0,
				Txs: []UserTransaction{
					&Transaction{
						TxID:      "tddxidddddd",
						IsDeposit: true,
						Amount:    10000,
					}, &Transaction{
						TxID:      "aaatxiddaaaa",
						IsDeposit: false,
						Amount:    40004,
					}},
				Plans: []UserPlan{
					&Plan{
						Start:    time.Now(),
						Active:   true,
						Type:     Type2,
						Invested: 100000,
					}, &Plan{
						Start:    time.Now(),
						Active:   true,
						Type:     Type3,
						Invested: 100000,
					}},
			}}, args: args{id: 1},
			want: &User{
				Id:             1,
				Name:           "Walter",
				DepositAddress: "ssssss",
				ReceiveAddress: "xxxxxxxx",
				ParentId:       0,
				Txs: []UserTransaction{
					&Transaction{
						TxID:      "tddxidddddd",
						IsDeposit: true,
						Amount:    10000,
					}, &Transaction{
						TxID:      "aaatxiddaaaa",
						IsDeposit: false,
						Amount:    40004,
					}},
				Plans: []UserPlan{
					&Plan{
						Start:    time.Now(),
						Active:   true,
						Type:     Type2,
						Invested: 100000,
					}, &Plan{
						Start:    time.Now(),
						Active:   true,
						Type:     Type3,
						Invested: 100000,
					}},
			},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if _, err := data.Insert(tt.fields.u); err != nil {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := data.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got.GetID(), tt.want.GetID())
			assert.Equal(t, got.GetDepositAddress(), tt.want.GetDepositAddress())
			assert.Equal(t, got.GetReceiveAddress(), tt.want.GetReceiveAddress())
			assert.Equal(t, got.GetParentId(), tt.want.GetParentId())
			assert.Equal(t, got.GetActivePlans()[0].GetStartDate().Day(), tt.want.GetActivePlans()[0].GetStartDate().Day())
			assert.Equal(t, got.GetDepositTransaction(), tt.want.GetDepositTransaction())
			assert.Equal(t, got.GetReceiveTransaction(), tt.want.GetReceiveTransaction())

		})
	}
}

func TestData_Insert(t *testing.T) {
	type args struct {
		u BotUser
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{name: "ok",
			args: args{&User{
				Id:             123,
				Name:           "maria",
				DepositAddress: "qweqweqwe",
				ReceiveAddress: "ewqewqewq",
				Balance:        "12345",
				ParentId:       0,
				Txs: []UserTransaction{
					&Transaction{
						TxID:      "txidddddd",
						IsDeposit: true,
						Amount:    10000,
					}, &Transaction{
						TxID:      "txiddaaaa",
						IsDeposit: false,
						Amount:    40004,
					}},
				Plans: []UserPlan{
					&Plan{
						Start:    time.Now(),
						Active:   true,
						Type:     Type1,
						Invested: 100000,
					}},
			}},
			want: 123, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := data.Insert(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Insert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestData_List(t *testing.T) {
// 	type fields struct {
// 		Db    *sql.DB
// 		Stmts map[string]*stmtConfig
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		want    []BotUser
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := Data{
// 				Db:    tt.fields.Db,
// 				Stmts: tt.fields.Stmts,
// 			}
// 			got, err := d.List()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("List() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestData_Update(t *testing.T) {
// 	type fields struct {
// 		Db    *sql.DB
// 		Stmts map[string]*stmtConfig
// 	}
// 	type args struct {
// 		id   int64
// 		user BotUser
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			d := Data{
// 				Db:    tt.fields.Db,
// 				Stmts: tt.fields.Stmts,
// 			}
// 			if err := d.Update(tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
// 				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
