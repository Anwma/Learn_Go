package main

import "testing"

func TrailingZeros(n int64) int64 {
	// write your code here
	res := int64(0)
	for n > 0 {
		n /= 5
		res += n
	}
	return res
}

func BenchmarkTrailingZeros(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrailingZeros(3)
	}
}