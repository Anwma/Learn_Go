package main

import "fmt"

func main() {
	//0 0 0 0 0
	//for i := 0; i < 5; i++ {
	//	//每次都重新申请 变量 v
	//	var v int
	//	fmt.Printf("%d ", v)
	//	v = 5
	//}
	//Value of i is now:0
	//Value of i is now:1 ... 无限循环
	//for i := 0; ; i++ {
	//	fmt.Println("Value of i is now:", i)
	//}
	//无线循环 0
	//for i := 0; i < 3; {
	//	fmt.Println("Value of i:", i)
	//}
	//"" a aa aaa aaaa
	//s := ""
	//for s != "aaaaa" {
	//	fmt.Println("Value of s:", s)
	//	s = s + "a"
	//}
	//0 5 "a"
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
	//0 5 a
	//1 6 aa
	//2 7 aaa
}
