package utils

import (
	"testing"
)

func TestCheckHashPassword(t *testing.T) {
	type args struct {
		hash     string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "admin", args: args{hash: HashPassword("admin"), password: "admin"}, want: true},
		{name: "user1", args: args{hash: HashPassword("user1"), password: "user1"}, want: true},
		{name: "user2", args: args{hash: HashPassword("user2"), password: "user2"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashPasswordValid(tt.args.hash, tt.args.password); got != tt.want {
				t.Errorf("HashPasswordValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
