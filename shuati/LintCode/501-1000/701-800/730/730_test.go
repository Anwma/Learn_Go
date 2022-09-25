package main

import (
	"testing"
)

func SubSum(n int) int {
	// write your code here
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	//return sum * int(math.Pow(float64(2), float64(n-1)))
	return sum * pow(2, n-1)
}
func pow(x, y int) int {
	for i := 0; i < y; i++ {
		x = x ^ x + (x&x)<<1
	}
	return x
}
func BenchmarkSubSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SubSum(100)
	}
}
