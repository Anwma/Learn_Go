package main

import "fmt"

var arr = [50]int{1, 1}

//1 1 2 3 5 8 13 21 34 55
func main() {
	for i := 1; i < 50; i++ {
		fmt.Println(fb(i))
	}
}
func fb(n int) (res int) {
	res = n
	if n <= 2 {
		return 1
	}
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
		res = arr[i]
	}
	return res
}
