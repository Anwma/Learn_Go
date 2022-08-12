package _14

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
	nums := []float32{1.0, 2.1, -3.3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxOfArray(nums)
	}
	b.StopTimer()
}
