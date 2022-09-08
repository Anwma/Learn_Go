package main

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
	for i := 0; i < b.N; i++ {
		Calculate(2, '+', 3)
	}
}
