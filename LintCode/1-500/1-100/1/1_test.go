package _1

import "testing"

func Aplusb(a int, b int) int {
	// write your code here
	return a+b
}
func BenchmarkAplusb(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Aplusb(2,3)
	}
	b.StopTimer()
}