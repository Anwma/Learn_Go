package main

func main() {

}

func twoSum(nums []int, target int) []int {
	// T = O(n^2) S = O(1)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func twoSum1(nums []int, target int) []int {
	// T = O(n) S = O(n)
	m := make(map[int]int)
	for k, v := range nums {
		other := target - v
		if index, ok := m[other]; ok {
			return []int{index, k}
		}
		m[v] = k
	}
	return []int{-1, -1}
}
