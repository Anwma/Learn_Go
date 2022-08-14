package _61

import "sort"

func KthSmallest(k int, nums []int) int {
	// write your code here
	sort.Ints(nums)
	return nums[k-1]
}
