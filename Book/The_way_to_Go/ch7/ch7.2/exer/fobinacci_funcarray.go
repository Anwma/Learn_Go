package main

import "fmt"

func main() {
	//	主函数调用一个使用序列个数 n 作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片。
	fmt.Println(fb(8))
}
func fb(n int) []int {
	//建立一个大小为n的切片 slice
	slice := make([]int, n)
	slice[0] = 1
	slice[1] = 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice
}
