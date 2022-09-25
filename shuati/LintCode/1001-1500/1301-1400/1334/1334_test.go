package _334

import "testing"

// 1.利用append特性 38.08 ns/op
func Rotate(nums []int, k int) []int {
	// Write your code here
	n := len(nums)
	return append(nums[n-k%n:], nums[:n-k%n]...)
}

// 2.反转三次数组 17.00 ns/op
func Rotate1(nums []int, k int) []int {
	l := len(nums)
	k = k % l
	if k == 0 {
		return nums
	}
	reverse(nums, 0, l-k-1)
	reverse(nums, l-k, l-1)
	reverse(nums, 0, l-1)
	return nums
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

// 3.新数组
func Rotate2(nums []int, k int) []int {
	n := len(nums)
	newArr := make([]int, n)
	for i := 0; i < n; i++ {
		newArr[(i+k)%n] = nums[i]
	}
	return newArr
}

func BenchmarkRotate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Rotate2(nums, 3)
	}
	b.StopTimer()
}
