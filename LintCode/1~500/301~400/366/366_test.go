package _66

import "testing"

func Fibonacci(n int) int {
	// write your code here
	q, p := 0, 1
	if n <= 2 {
		return n - 1
	}
	for i := 2; i < n; i++ {
		sum := q + p
		q = p
		p = sum
	}
	return p
}

func BenchmarkFibonacci(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibonacci(1)
	}
	b.StopTimer()
}
