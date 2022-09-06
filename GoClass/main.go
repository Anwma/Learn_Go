package main

import (
	"GoClass/dog"
	"fmt"
)
import . "GoClass/testpackage"

func main() {
	var a string = "1010"
	//关键字 变量名 变量类型 = "变量值" //作为注释

	/*
		无论我写什么，系统都看不到
	*/
	b := "2020"
	fmt.Print("hello " + a)
	fmt.Print("hello " + b)
	fmt.Print(A)
	fmt.Print(dog.Name)
	str := "123"
	fmt.Printf("%T",str)
	aa := 1.23
	fmt.Print(int(aa))
	_aaa := 1
	fmt.Print(_aaa)

}
