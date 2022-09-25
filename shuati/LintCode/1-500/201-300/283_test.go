package main

import "testing"

func MaxOfThreeNumbers(num1 int, num2 int, num3 int) int {
	// write your code here
	return max(num3, max(num1, num2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func BenchmarkMaxOfThreeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxOfThreeNumbers(1, 2, 3)
	}
}
