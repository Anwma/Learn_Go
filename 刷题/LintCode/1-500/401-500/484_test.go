package main

import "testing"

func SwapIntegers(a []int, index1 int, index2 int) {
	// write your code here
	a[index1], a[index2] = a[index2], a[index1]
}

func BenchmarkSwapIntegers(b *testing.B) {
	a := []int{1, 2}
	for i := 0; i < b.N; i++ {
		SwapIntegers(a, 0, 1)
	}
}
