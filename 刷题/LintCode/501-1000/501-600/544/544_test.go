package main

/**
 * @param nums: an integer array
 * @param k: An integer
 * @return: the top k largest numbers in array
 */
import "sort"

func Topk(nums []int, k int) []int {
	// write your code here
	sort.Ints(nums)
	swap(nums, 0, len(nums)-1)
	return nums[:k]
}

func swap(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
