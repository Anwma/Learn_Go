package main

import "fmt"

func ProductExcludeItself(nums []int) []int64 {
	// write your code here
	res := make([]int64, len(nums))
	for i := 0; i < len(nums); i++ { //1-n
		tmp := 1
		for j := 0; j < len(nums); j++ { //1-n
			if i != j { //去除本身的值
				tmp *= nums[j]
			}
		}
		res[i] = int64(tmp)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(ProductExcludeItself(nums))
}
