package _7

import "testing"

func ReverseInteger(number int) int {
	// write your code here
	return number%10*100 + number/10%10*10 + number/100
}
func BenchmarkReverseInteger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseInteger(213)
	}
	b.StopTimer()
}