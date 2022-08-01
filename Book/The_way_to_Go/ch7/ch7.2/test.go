package main

import "fmt"

func main() {
	arr1 := make([]int, 50, 100)
	//new 后面的参数指定 start 和 end(len) 。若添加第三个参数则指定cap
	arr2 := new([100]int)[0:30]
	fmt.Println(len(arr1), cap(arr1))
	fmt.Println(len(arr2), cap(arr2))
}
