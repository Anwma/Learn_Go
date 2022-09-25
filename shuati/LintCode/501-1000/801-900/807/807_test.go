package _07

import (
	"math"
	"testing"
)

func IsPalindrome(n int) bool {
	// Write your code here
	if n <= 1 {
		return true
	}
	a := toBinary(n)
	for i := 0; i <= len(a)-i-1; i++ {
		if a[i] == a[len(a)-i-1] {
			continue
		} else {
			return false
		}
	}
	return true
}

func toBinary(n int) []int {
	a := make([]int, 0, 32)
	for {
		a = append(a, n%2)
		n = n / 2
		if n == 1 {
			a = append(a, 1)
			break
		}
	}
	return a
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(math.MaxInt16 - 1)
	}
	b.StopTimer()
}
