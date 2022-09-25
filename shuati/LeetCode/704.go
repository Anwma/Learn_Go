package main

//顺序查找
//func search(nums []int, target int) int {
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == target {
//			return i
//		}
//	}
//	return -1
//}

// 二分查找
func search(nums []int, target int) int {
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
	if nums[l] == target {
		return l
	}
	if nums[r] == target {
		return r
	}
	return -1
}
