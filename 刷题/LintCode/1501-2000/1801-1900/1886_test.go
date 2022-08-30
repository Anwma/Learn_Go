package main

import "testing"

// 5861778               187.4 ns/op
func MoveTarget(nums []int, target int) {
	// write your code here
	//1..创建两个切片 大于target 的放在右切片 小于等于放在左切片
	arr := make([]int, 0)
	brr := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			arr = append(arr, target)
		} else {
			brr = append(brr, nums[i])
		}
	}
	copy(nums, append(arr, brr...))
}

// 93437568                10.72 ns/op
func MoveTarget1(nums []int, target int) {
	// write your code here
	//1.左右两个指针 从尾部开始 left right 两指针
	count := 0
	left, right := len(nums)-1, len(nums)-1
	for left >= 0 {
		if nums[left] != target {
			nums[right] = nums[left]
			right -= 1
		} else {
			count++
		}
		left -= 1
	}
	for i := 0; i < count; i++ {
		nums[i] = target
	}
}

func BenchmarkMoveTarget(b *testing.B) {
	nums := []int{5, 1, 6, 1}
	for i := 0; i < b.N; i++ {
		MoveTarget(nums, 1)
	}
}
