package main

import "fmt"

func main() {
	//var a byte
	//a = 'a' //输出ascii对应码值 。。 这里说明一下什么是ascii码
	//fmt.Println(a)
	//fmt.Printf("a=%c", a)
	a := 300.0
	for i := 0; i < 30; i++ {
		a *= 1.01
	}
	fmt.Println(a)
}
