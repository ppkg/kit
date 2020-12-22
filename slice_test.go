package base_tool

import (
	"testing"
)

func TestInSlice(t *testing.T) {
	type args struct {
		value interface{}
		list  interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "in slice",
			args: args{
				value: 1,
				list:  []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "not in slice",
			args: args{
				value: "5",
				list:  []string{"1", "2", "3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSlice(tt.args.value, tt.args.list); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
