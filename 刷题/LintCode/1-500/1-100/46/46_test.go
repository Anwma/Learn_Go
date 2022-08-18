package main

import (
	"testing"
)

//方法一：排序好之后直接返回中间值
//20618131                56.69 ns/op           24 B/op          1 allocs/op

//func MajorityNumber(nums []int) int {
//	// write your code here
//	sort.Ints(nums)
//	return nums[len(nums)/2]
//}

//方法二：利用hashmap
//22833091                50.70 ns/op            0 B/op          0 allocs/op


//func MajorityNumber(nums []int) int {
//	// write your code here
//	count := len(nums)/2
//	m := make(map[int]int)
//	for _, num := range nums {
//		m[num]++
//		if m[num] > count{
//			return num
//		}
//	}
//	return -1
//}

// 方法三：Boyer-Moore 投票算法
//121280895                9.864 ns/op           0 B/op          0 allocs/op

func MajorityNumber(nums []int) int {
	// write your code here
	count, candidate := 0, -1
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count*2 > len(nums) {
		return candidate
	}
	return -1
}

func BenchmarkMajorityNumber(b *testing.B) {
	nums := []int{1, 1, 1, 1, 2, 2, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MajorityNumber(nums)
	}
	b.StopTimer()
}
