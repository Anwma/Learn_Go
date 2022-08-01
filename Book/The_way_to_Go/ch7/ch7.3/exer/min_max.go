package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	min := minSlice(slice)
	max := maxSlice(slice)
	fmt.Println("最小值:", min)
	fmt.Println("最大值:", max)
}

//传入一个int切片，返回最小值
func minSlice(s []int) int {
	min := s[0]
	for i := 1; i < len(s); i++ {
		if min > s[i] {
			min = s[i]
		}
	}
	return min
}
func maxSlice(s []int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if max < s[i] {
			max = s[i]
		}
	}
	return max
}
