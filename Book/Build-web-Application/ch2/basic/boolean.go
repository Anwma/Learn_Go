package main

import "fmt"

func main() {
	//示例代码
	var isActive bool                   // 全局变量声明
	var enabled, disabled = true, false // 忽略类型的声明
	fmt.Println(isActive)
	fmt.Println(enabled, disabled)

}

func test() {
	var available bool // 一般声明
	valid := false     // 简短声明
	fmt.Println(available, valid)
	available = true // 赋值操作
}
