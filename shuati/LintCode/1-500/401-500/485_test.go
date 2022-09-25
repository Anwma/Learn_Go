package main

import "testing"

func Generate(size int) []int {
	// write your code here
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i + 1
	}
	return slice
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(100)
	}
}
