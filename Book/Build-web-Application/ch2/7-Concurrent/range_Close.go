package main

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		//接收值 将 x 的值发送到channel上
		c <- x
		x, y = y, x+y
	}
	//for i := range 能不断从chan 中读取数据,所以此处需要close channel
	close(c)
}

func main() {
	//1. make一个chan,并指定容量
	c := make(chan int, 10)
	//2.起一个协程运行fib函数,fib函数接收两个参数：
	//因为要接收chan 中的值,参数有channel.又因为要求前n个数的fib。所以有n
	//① channel ② n
	go fibonacci(cap(c), c)
	//3.打印值, 从channel当中获取
	for i := range c {
		fmt.Println(i)
	}
}
