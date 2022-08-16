package main

func ProductExcludeItself(nums []int) []int64 {
	// write your code here
	res := make([]int64, len(nums))
	tmp := 1
	mul := 1
	for i := 0; i < len(nums); i++ { //1-n
		for j := 0; j < len(nums); j++ {//1-n
			if i != j {//去除本身的值
				mul = tmp * nums[j]
			}
		}
		res[i] = int64(mul)
	}
	return res
}


