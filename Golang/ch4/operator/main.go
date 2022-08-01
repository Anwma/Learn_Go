package main

import "fmt"

func main() {
	//算术运算符
	//a := 12
	//b := 22
	//fmt.Println(a - b)
	//fmt.Println(a + b)
	//fmt.Println(a * b)
	//fmt.Println(a / b)
	//a++
	//++a 不能前++
	var a = true
	var b = false
	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}
