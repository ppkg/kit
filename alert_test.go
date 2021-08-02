package kit

import "testing"

func Test_wecomAlert(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "企业微信警报",
			args: args{content: "test alert"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wecomAlert(tt.args.content)
		})
	}
}
