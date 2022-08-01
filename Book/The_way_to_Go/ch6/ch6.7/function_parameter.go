package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// callback(1,Add)
// 调用callback
//传递参数到 callback (y, f) y是一个值 = 1 f是一个接收两个值的函数f(int int) 将Add 参数传递给 f
//y=1-> a ; 2 -> b

func IsAscii(c int) bool {
	if c > 255 {
		return false
	}
	return true
}
