package main

import (
	"fmt"
)

func main() {
	var x int = 12
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(x); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	fmt.Println(x)

	var integer int = x - 9
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}
}

func computedValue(x1 int) (x int) {
	return x1 + x1
}
