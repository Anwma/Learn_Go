package main

import "fmt"

func main() {
	//1.创建array type.[3]int(SB), AX
	//2.新建一个slice{arr,len,cap}
	s := []int{1, 2, 3}
	fmt.Println(cap(s))
	s = append(s, 4, 5)
	fmt.Println(cap(s))

	//d := make([]int, 3)
	//fmt.Println(d)

	//reflect.SliceHeader{
	//	Data: 0,
	//	Len:  0,
	//	Cap:  0,
	//}
	//fmt.Println(s)
}
