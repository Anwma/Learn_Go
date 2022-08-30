package main

import "testing"

/**
 * @param nums: an array of integers
 * @param s: An integer
 * @return: an integer representing the minimum size of subarray
 */

func BenchmarkMinimumSize(b *testing.B) {
	s := []int{100, 50, 99, 50, 100, 50, 99, 50, 100, 50}
	for i := 0; i < b.N; i++ {
		MinimumSize(s, 250)
	}
}

//暴力 n^3
/*
func MinimumSize1(nums []int, s int) int {
	// write your code here
	l := len(nums)
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if s <= Sum(nums, i, j) {
				flag = true
				if l > j-i+1 {
					l = j - i + 1
				}
			}
		}
	}
	if flag == false {
		return -1
	}
	return l
}

func Sum(nums []int, left, right int) (res int) {
	for i := left; i <= right; i++ {
		res += nums[i]
	}
	return res
}
*/

func MinimumSize(nums []int, s int) int {
	// write your code here
	sum, res := 0, len(nums)+1
	left, right := 0, 0
	for right < len(nums) {
		//右指针
		for sum < s && right < len(nums) {
			sum += nums[right]
			right++
		}
		for sum >= s {
			if res > right-left {
				res = right - left
			}
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return -1
	}
	return res
}
