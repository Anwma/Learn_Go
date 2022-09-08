package main

import (
	"fmt"
	"strings"
	"testing"
)

func DigitCounts(k int, n int) int {
	// write your code here
	count := 0
	strK := fmt.Sprintf("%d", k)
	for i := 0; i <= n; i++ {
		count += len(strings.Split(fmt.Sprintf("%d", i), strK)) - 1
	}
	return count
}

func BenchmarkDigitCounts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitCounts(5, 3)
	}
}
