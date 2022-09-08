package main

import "testing"

func BinarySearch(nums []int, target int) int {
	// write your code here
	if nums == nil {
		return -1
	}
	l, r := 0, len(nums)-1
	for l+1 < r {
		m := l + (r-l)>>1
		//要求第一次出现下表 所以先把target <= nums[m] 先写
		if target <= nums[m] {
			r = m
		} else {
			l = m
		}
	}
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func BenchmarkBinarySearch(b *testing.B) {
	nums := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		BinarySearch(nums, 2)
	}
}
