package main

import "testing"

func MaxOfArray(a []float32) float32 {
	// write your code here
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	return max
}
func BenchmarkMaxOfArray(b *testing.B) {
	a := []float32{1, 2, 4, 9}
	for i := 0; i < b.N; i++ {
		MaxOfArray(a)
	}
}
