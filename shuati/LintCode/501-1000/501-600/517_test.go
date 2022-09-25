package main

import "testing"

//func IsUgly(num int) bool {
//	write your code here
//if num <= 0 {
//	return false
//}
//for {
//	if num%2 == 0 {
//		num /= 2
//	} else if num%3 == 0 {
//		num /= 3
//	} else if num%5 == 0 {
//		num /= 5
//	} else if num == 1 {
//		return true
//	} else {
//		return false
//	}
//}
//}

var factors = []int{2, 3, 5}

func IsUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}

func BenchmarkIsUgly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsUgly(15)
	}
}
