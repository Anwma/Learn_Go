package main

import "C"
import "fmt"

// 导出Go函数声明, 给C使用
//
//export GoFunction
func GoFunction() {
	fmt.Println("GoFunction!!!")
}

