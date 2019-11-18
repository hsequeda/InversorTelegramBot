package main

import (
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

// func TestData_Get(t *testing.T) {
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
// 		want    BotUser
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
// 			got, err := d.Get(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Get() got = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

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
