package main

import (
	"strconv"
	"testing"
)

func FizzBuzz(n int) []string {
	if n == 0 {
		return []string{}
	}
	str := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			str[i-1] = "fizz buzz"
		case i%5 == 0:
			str[i-1] = "buzz"
		case i%3 == 0:
			str[i-1] = "fizz"
		default:
			str[i-1] = strconv.Itoa(i)
		}
	}
	return str
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBuzz(8)
	}
}
