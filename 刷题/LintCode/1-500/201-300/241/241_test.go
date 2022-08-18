package _41

import (
	"strconv"
	"testing"
)

func StringToInteger(target string) int {
	// write your code here
	v, _ := strconv.Atoi(target)
	return v
}

func BenchmarkStringToInteger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToInteger("123")
	}
	b.StopTimer()
}
