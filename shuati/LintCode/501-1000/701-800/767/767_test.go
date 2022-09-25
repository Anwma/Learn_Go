package _67

import "testing"

func ReverseArray(nums []int) {
	// write your code here
	n := len(nums) - 1
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[n-i] = nums[n-i], nums[i]
	}
}

func BenchmarkReverseArray(b *testing.B) {
	nums := []int{1, 2, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseArray(nums)
	}
	b.StopTimer()
}
