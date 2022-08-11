package _141

import "testing"

func GetTheMonthDays(year int, month int) int {
	// write your code here
	if month == 2 {
		if IsLeapYear(year) {
			return 29
		}
		return 28
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}
func IsLeapYear(n int) bool {
	// write your code here
	if n%4 == 0 && n%100 != 0 || n%400 == 0 {
		return true
	}
	return false
}
func BenchmarkGetTheMonthDays(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetTheMonthDays(2022, 8)
	}
	b.StopTimer()
}
