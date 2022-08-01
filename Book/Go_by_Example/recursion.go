package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	//计算7的阶乘
	fmt.Println(fact(7))
	//1 2 6 24 120 720 5040...
	var fib func(n int) int
	//计算斐波那契数列
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)

	}
	//1 1 2 3 5 8 13...
	fmt.Println(fib(7))
}
