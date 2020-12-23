package rp_kit

import (
	"testing"
	"time"
)

func Test_GetTimeNow(t *testing.T) {
	type args struct {
		time2 []time.Time
	}
	time2, _ := time.Parse(DATETIME_LAYOUT, "2020-11-25 00:00:00")
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "获取当前格式化时间",
			args: args{},
			want: time.Now().Format(DATETIME_LAYOUT),
		},
		{
			name: "获取指定格式化时间",
			args: args{
				time2: []time.Time{time2},
			},
			want: "2020-11-25 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeNow(tt.args.time2...); got != tt.want {
				t.Errorf("GetTimeNow() = %v, want %v", got, tt.want)
			}
		})
	}
}
