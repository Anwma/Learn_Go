package main

import "fmt"

func main() {
	var src = []int{0, 1, 2, 3, 4, 5}
	fmt.Println(src)
	//左取右也取
	src = RemoveStringSlice(src, 2, 3)
	fmt.Println(src)
}

// RemoveStringSlice 传入源slice start end
func RemoveStringSlice(src []int, start int, end int) []int {
	slice := make([]int, 0)
	slice = append(append(slice, src[0:start]...), src[end+1:]...)
	return slice
}
