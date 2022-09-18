package main

import "fmt"

func main() {
	inputs := []int{-3, 5, 3, 2, 8}
	tests := []int{-1, 0, 1, 2}
	res := AddAndSearch(inputs, tests)
	fmt.Print(res)
}

func AddAndSearch(inputs []int, tests []int) bool {
	// write your code here.
	n := len(inputs)
	//要从inputs中找到一对，所以需要数组长度不小于2 也就是 >= 2
	if n < 2 {
		return false
	}
	res := make([]int, (n*n-n)/2)
	index := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			res[index] = inputs[i] + inputs[j]
			index++
		}
	}
	for i := 0; i < len(tests); i++ {
		for j := 0; j < len(res); j++ {
			if tests[i] == res[j] {
				return true
			}
		}
	}
	return false
}
