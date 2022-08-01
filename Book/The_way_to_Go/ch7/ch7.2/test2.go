package main

import "fmt"

func main() {
	//类型为 *[]int 指向 [ptr,len,cap] = [nil, 0 , 0] 的一个 [] int
	arr1 := new([]int)
	fmt.Printf("%T\n", arr1)
	//类型为 []int [ptr,len,cap] = [[0]int,0,0]
	arr2 := make([]int, 0, 0)
	fmt.Printf("%T\n", arr2)

}
