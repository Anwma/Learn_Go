package _21

import (
	"sort"
	"testing"
)

func Deduplication(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	n := len(nums)
	i, j := 0, 1
	for ; i < n; i++ {
		for j < n && nums[j] == nums[i] {
			j++
		}

		if j >= n {
			break
		}
		nums[i+1] = nums[j]
	}
	return i + 1
}
func BenchmarkDeduplication(b *testing.B) {
	nums := []int{1, 3, 1, 4, 4, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Deduplication(nums)
	}
	b.StopTimer()
}
