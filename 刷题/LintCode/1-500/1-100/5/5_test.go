package main

import "sort"

func KthLargestElement(k int, nums []int) int {
	// write your code here
	sort.Ints(nums)
	return nums[len(nums)-k]
}
