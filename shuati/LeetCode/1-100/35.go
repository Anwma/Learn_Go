package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l+1 < r {
		m := l + (r-l)>>1
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			r = m
		} else {
			l = m
		}
	}
	if nums[l] >= target {
		return l
	} else if nums[r] >= target {
		return r
	} else if nums[r] < target {
		return r + 1
	} else {
		return l + 1
	}
}
