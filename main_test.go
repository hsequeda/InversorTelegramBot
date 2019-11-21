package main

import "testing"

func Test_getParentIdFromMessage(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{{name: "ok", args: args{msg: "/start 1234"}, want: 1234}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getParentIdFromMessage(tt.args.msg); got != tt.want {
				t.Errorf("getParentIdFromMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
