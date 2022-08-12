package _83

import "testing"

func MaxOfThreeNumbers(num1 int, num2 int, num3 int) int {
	// write your code here
	tmp := max(num1, num2)
	return max(tmp, num3)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func BenchmarkMaxOfThreeNumbers(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxOfThreeNumbers(1, 2, 4)
	}
	b.StopTimer()
}
