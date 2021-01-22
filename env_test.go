package rp_kit

import (
	"testing"
)

func TestEnvExistsValue(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			// ProgramData=C:\ProgramData
			name: "test1",
			args: args{
				name:  "ProgramData",
				value: `C:\ProgramData`,
			},
			want: true,
		},
		{
			// ProgramData=C:\ProgramData
			name: "test2",
			args: args{
				name:  "ProgramData",
				value: `FALSE`,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EnvExistsValue(tt.args.name, tt.args.value); got != tt.want {
				t.Errorf("EnvExistsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
