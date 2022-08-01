package main

import "fmt"

func main() {
	//定义一个名称为“variableName”，类型为 “type”的变量
	var variableName int
	fmt.Println(variableName)
	//定义三个类型都是“type”的变量
	var vname1, vname2, vname3 int
	fmt.Println(vname1, vname2, vname3)

	//初始化“variableName”的变量为“value”值，类型是“type”
	var variableName2 = 3
	fmt.Println(variableName2)

	/*
	   定义三个类型都是"type"的变量,并且分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	*/
	var vname_1, vname_2, vname_3 = 1, 2, 3
	fmt.Println(vname_1, vname_2, vname_3)

	/*
	   定义三个变量，它们分别初始化为相应的值
	   vname1为v1，vname2为v2，vname3为v3
	   编译器会根据初始化的值自动推导出相应的类型
	*/
	vname__1, vname__2, vname__3 := 1, 2, 3
	fmt.Println(vname__1, vname__2, vname__3)

	_, b := 34, 35
	fmt.Println(b)

	var i int
	fmt.Println(i)
}
