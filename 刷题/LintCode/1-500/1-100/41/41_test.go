package _1

/**
 * @param nums: A list of integers
 * @return: A integer indicate the sum of max subarray
 */

func maxSubArray (nums []int) int {
	if len(nums) == 1 {return nums[0]}
	Inf := 2147483647
	total := 0
	maxSum := nums[0]
	minSum := Inf
	for i := 0; i < len(nums); i++ {
		total += nums[i]
		maxSum = max(maxSum,total - minSum)
		maxSum = max(maxSum,total)
		minSum = min(minSum,total)
	}
	return maxSum
}
func max(a,b int) int {
	if a > b {return a}
	return b
}
func min(a,b int) int {
	if a < b {return a}
	return b
}
