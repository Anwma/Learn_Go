package _334

import "testing"

func Rotate(nums []int, k int) []int {
	// Write your code here
	n := len(nums)
	return append(nums[n-k%n:], nums[:n-k%n]...)

	//l := len(nums)
	//if k%l==0 {
	//	return nums
	//}
	//k = k%l
	//reverse(nums,0,l-k-1)
	//reverse(nums,l-k,l-1)
	//reverse(nums,0,l-1)
	//return nums

	//newArr := make([]int,n)
	//for i := 0; i < n; i++ {
	//	newArr[(i+k)%n] = nums[i]
	//}
	//return newArr
}
func gcd(x, y int) int {
	if y > 0 {
		return gcd(y, x%y)
	} else {
		return x
	}
}

func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func BenchmarkRotate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Rotate(nums, 3)
	}
	b.StopTimer()
}
