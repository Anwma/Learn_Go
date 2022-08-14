package main

import (
	"sort"
	"testing"
)

//解法一:

func RecoverRotatedSortedArray(nums []int) {
	//write your code here
	//BenchmarkName-8         24183114                51.44 ns/op
	sort.Ints(nums)
}

// 解法二: 迭代
//func RecoverRotatedSortedArray(nums []int) {
//	//write your code here
//	//BenchmarkName-8         190293206                6.320 ns/op
//	//1.先找到后面一个数比前面大的索引值
//	index := 0
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] > nums[i+1] {
//			index = i
//			break
//		}
//	}
//	//当index不发生改变的时候，说明传入的[]int是顺序的,无需修改。
//	if index != 0 {
//		reverse(nums, 0, index)
//		reverse(nums, index+1, len(nums)-1)
//		reverse(nums, 0, len(nums)-1)
//	}
//}
//
//func reverse(nums []int, left, right int) {
//	for left < right {
//		nums[left], nums[right] = nums[right], nums[left]
//		left++
//		right--
//	}
//}

func BenchmarkName(b *testing.B) {
	num := []int{4, 5, 1, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RecoverRotatedSortedArray(num)
	}
	b.StopTimer()
}
