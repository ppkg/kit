package kit

import "testing"

func Test_YuanToFen(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				f: 4.2,
			},
			want: 420,
		},
		{
			name: "test2",
			args: args{
				f: 2.01,
			},
			want: 201,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := YuanToFen(tt.args.f); got != tt.want {
				t.Errorf("YuanToFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_FenToYuan(t *testing.T) {
	type args struct {
		f int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test1",
			args: args{
				f: 420,
			},
			want: 4.2,
		},
		{
			name: "test2",
			args: args{
				f: 201,
			},
			want: 2.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FenToYuan(tt.args.f); got != tt.want {
				t.Errorf("YuanToFen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Float32ToFloat64(t *testing.T) {
	type args struct {
		f float32
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Float32ToFloat64",
			args: args{
				f: 2.01,
			},
			want: 2.01,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32ToFloat64(tt.args.f); got != tt.want {
				t.Errorf("Float32ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
