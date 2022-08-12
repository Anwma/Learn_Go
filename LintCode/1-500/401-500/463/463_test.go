package _63

import (
	"sort"
	"testing"
)

func SortIntegers(a []int) {
	// write your code here
	sort.Ints(a)
}

func BenchmarkSortIntegers(b *testing.B) {
	a := []int{1, 5, 8, 2, 4, 9}
	for i := 0; i < b.N; i++ {
		SortIntegers(a)
	}
}
