package main

func MaxSlidingWindow(nums []int, k int) []int {
	// write your code here
	//1.corner case
	if len(nums) == 0{
		return []int{}
	}
	max := nums[0]
	res := make([]int, len(nums)-k+1)
	//先计算前n个数 找到第一个 max 值
	for i := 1; i < k; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	res[0] = max
	for i := k; i < len(nums); i++ {
		res[i-k+1] = Max(max, nums[i])
	}
	return res
}

func Max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}
