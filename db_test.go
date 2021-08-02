package kit

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"github.com/limitedlee/microservice/common/config"
)

func Test_NewRedisEngine(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "创建redis连接",
			args: args{dsn: config.GetString("redis.Addr") + "|" + config.GetString("redis.Password") + "|" + config.GetString("redis.DB")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRedisEngine(tt.args.dsn)
			if err := got.Engine.(*redis.Client).Ping().Err(); err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_NewDBEngine(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "创建mysql连接",
			args: args{dsn: config.GetString("mysql.test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDBEngine(tt.args.dsn)
			if err := got.Engine.(*xorm.Engine).Ping(); err != nil {
				t.Error(err)
			}
		})
	}
}
