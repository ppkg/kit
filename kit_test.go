package rp_kit

import "testing"

func Benchmark_GetGuid32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetGuid32()
	}
}

func Benchmark_RunFuncName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunFuncName()
	}
}

func Benchmark_GetMd5String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetMd5("9d9dce8ec1654ee28ad50ede7e04247b")
	}
}

func Benchmark_GetRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetRandomString(i)
	}
}

func Test_UrlDecode(t *testing.T) {
	type args struct {
		org string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "url字符解码",
			args: args{
				org: `+%E3%80%90%E5%A6%82%E9%81%87%E7%BC%BA%E8%B4%A7%E3%80%91%EF%BC%9A+%E7%BC%BA%E8%B4%A7%E6%97%B6%E7%94%B5%E8%AF%9D%E4%B8%8E%E6%88%91%E6%B2%9F%E9%80%9A+%E9%A1%BE%E5%AE%A2%E6%9C%AA%E5%AF%B9%E9%A4%90%E5%85%B7%E6%95%B0%E9%87%8F%E5%81%9A%E9%80%89%E6%8B%A9`,
			},
			want: "+【如遇缺货】：+缺货时电话与我沟通+顾客未对餐具数量做选择",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UrlDecode(tt.args.org); got != tt.want {
				t.Errorf("UrlDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_RunFuncName(t *testing.T) {
	type args struct {
		skip []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "获取当前函数名",
			args: args{},
			want: "github.com/legofun/go-pkg.TestRunFuncName.func1",
		},
		{
			name: "获取上层调用函数名",
			args: args{skip: []int{2}},
			want: "testing.tRunner",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunFuncName(tt.args.skip...); got != tt.want {
				t.Errorf("RunFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
