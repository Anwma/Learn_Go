package _97

import (
	"sort"
	"testing"
)

func MaxNum(nums []int) int {
	// write your code here
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func BenchmarkMaxNum(b *testing.B) {
	n := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxNum(n)
	}
	b.StopTimer()
}
