package main

func main() {
	nums := []int{1, 2, 7, 7, 2, 10, 3, 4, 5}
	MaxSlidingWindow(nums, 2)
}

//暴力

//func MaxSlidingWindow(nums []int, k int) []int {
//	// write your code here
//	//1.corner case
//	if len(nums) == 0 {
//		return []int{}
//	}
//	res := make([]int, len(nums)-k+1)
//	for i := 0; i < len(nums)-k+1; i++ {
//		max := nums[i]
//		for j := i; j < i+k; j++ {
//			if nums[j] > max {
//				max = nums[j]
//			}
//		}
//		res[i] = max
//	}
//	return res
//}

//大顶堆优先队列

//双端队列


