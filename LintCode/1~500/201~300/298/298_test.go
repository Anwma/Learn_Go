package _98

import "testing"

func Prime(n int) []int {
	// write your code here
	//var slice []int
	//这个不行 会返回null 我自己本地1.18.4测试的ok,估计lintcode版本低了...
	slice := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			slice = append(slice, i)
		}
	}
	return slice
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func BenchmarkPrime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Prime(100)
	}
	b.StopTimer()
}
