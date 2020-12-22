package cast

import (
	"strconv"
	"testing"
)

func Benchmark_ToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToString(i)
	}
}

func Benchmark_ToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		ToInt(s)
	}
}
