package _300

import "testing"

func CanWinBash(n int) bool {
	// Write your code here
	if n % 4 == 0{
		return false
	}
	return true
}

func BenchmarkCanWinBash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanWinBash(4001)
	}
	b.StopTimer()
}
