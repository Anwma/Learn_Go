package _64

import "testing"

func Calculate(r int) []float64 {
	// write your code here
	PI := 3.14
	return []float64{float64(2*r) * PI, PI * float64(r*r)}
}

func BenchmarkCalculate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
	b.StopTimer()
}
