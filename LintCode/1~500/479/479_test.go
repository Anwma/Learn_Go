package _79

import (
	"sort"
	"testing"
)

func SecondMax(nums []int) int {
	// write your code here
	sort.Ints(nums)
	return nums[len(nums)-2]
}
func BenchmarkSecondMax(b *testing.B) {
	nums := []int{1,2,2,3,1,1,3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecondMax(nums)
	}
	b.StopTimer()
}
