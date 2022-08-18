package _49

import "testing"

func CharToInteger(character byte) int {
	// write your code here
	return int(character)
}

func BenchmarkCharToInteger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CharToInteger('a')
	}
	b.StopTimer()
}