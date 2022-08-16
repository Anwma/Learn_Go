package main

import "testing"

func DoubleFactorial(n int) int64 {
	// Write your code here
	var res int64 = 1
	if n%2 == 0 {
		for i := 2; i <=n ; i+=2 {
			res *= int64(i)
		}
	}else {
		for i := 1; i <= n; i+=2 {
			res *= int64(i)
		}
	}
	return res
}

func BenchmarkDoubleFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoubleFactorial(7)
	}
}
