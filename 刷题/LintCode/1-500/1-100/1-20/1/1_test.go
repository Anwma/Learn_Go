package main

import "testing"

func Aplusb(a int, b int) int {
	// write your code here
	return a + b
}

func BenchmarkAplusb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Aplusb(1, 2)
	}
}
