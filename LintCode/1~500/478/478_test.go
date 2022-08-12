package _78

import "testing"

func Calculate(a int, op byte, b int) int {
	// write your code here
	res := 0
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	}
	return res
}
func BenchmarkCalculate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate(5, '-', 3)
	}
	b.StopTimer()
}
