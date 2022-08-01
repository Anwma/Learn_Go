package main

import "fmt"

func main() {
	var f = Adder2()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

//在多次调用中，变量 x 的值是被保留的
//0 + 1 = 1
//1 + 20 = 21
//21 + 300 = 321

func Adder2() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
