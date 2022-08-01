package main

import "fmt"

func main() {
	var s = []int{1, 2, 3}
	var data = []int{4, 5}
	newSlice := InsertStringSlice(s, data, 3)
	fmt.Println(newSlice)
}

// InsertStringSlice 将切片插入到另一个切片的指定位置。
func InsertStringSlice(src []int, data []int, position int) []int {
	//新建一个容量为 len(src) + len(dst) 的 slice
	dst := make([]int, 0)
	dst = append(append(append(dst, src[0:position]...), data...), src[position:]...)
	return dst
}
