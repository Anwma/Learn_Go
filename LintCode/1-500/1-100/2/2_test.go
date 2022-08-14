package main

import "testing"

func TrailingZeros(n int64) int64 {
	// write your code here
	ans := int64(0)
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}

func BenchmarkTrailingZeros(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TrailingZeros(100)
	}
	b.StopTimer()
}