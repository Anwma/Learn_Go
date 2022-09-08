package main

import (
	"sort"
	"testing"
)

func SortIntegers(a []int) {
	// write your code here
	sort.Ints(a)
}

func BenchmarkSortIntegers(b *testing.B) {
	a := []int{3, 2, 1, 4, 5}
	for i := 0; i < b.N; i++ {
		SortIntegers(a)
	}
}
