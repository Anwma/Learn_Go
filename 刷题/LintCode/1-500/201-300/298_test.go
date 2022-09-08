package main

import "testing"

func Prime(n int) []int {
	// write your code here
	res := []int{}
	if n == 1 {
		return res
	}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func isPrime(num int) bool {
	for i := 2; i <= num/i; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func BenchmarkPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Prime(2022)
	}
}
