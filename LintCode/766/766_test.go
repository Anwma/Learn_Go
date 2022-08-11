package _66

import "testing"

func IsLeapYear(n int) bool {
	// write your code here
	if n%4 == 0 && n%100 != 0 || n%400 == 0 {
		return true
	}
	return false
}

func BenchmarkIsLeapYear(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsLeapYear(2022)
	}
	b.StopTimer()
}
