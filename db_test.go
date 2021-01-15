package rp_kit

import (
	"testing"
)

func Test_GetRedisConn(t *testing.T) {
	type args struct {
		dsn []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "GetRedisConn",
			args: args{
				dsn: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetRedisConn(tt.args.dsn...).Ping().Err(); err != nil {
				t.Error(err)
			}
		})
	}
}
