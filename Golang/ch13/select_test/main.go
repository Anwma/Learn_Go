package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		go语言提供了select的功能，作用与channel之上的 多路复用
		select 会随机公平的选择一个case语句执行
		select的应用场景 1.timeout的超时机制 2.判断某个channel是否阻塞
	*/
	//timeout := false
	timeout := make(chan bool, 2)
	go func() {
		//该goroutine如果执行时间超过了5s 那么就报告给主的goroutine
		//fmt.Println("结束")
		time.Sleep(time.Second * 5)
		timeout <- true
	}()
	timeout2 := make(chan bool, 2)
	go func() {
		//该goroutine如果执行时间超过了1s 那么就报告给主的goroutine
		//fmt.Println("结束")
		time.Sleep(time.Second * 1)
		timeout2 <- true
	}()
	select {
	case <-timeout:
		fmt.Println("超时了1")
	case <-timeout2:
		fmt.Println("超时了2")
	default:
		fmt.Println("继续下一次")
	}
	//fmt.Println(<-timeout)

	//for {
	//	if timeout {
	//		fmt.Println("结束")
	//		break
	//	}
	//	time.Sleep(time.Millisecond * 10)
	//}

	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//ch1 <- 1
	//ch2 <- 2
	////随机选择一个
	//select {
	//case data := <-ch1:
	//	fmt.Println(data)
	//case data := <-ch2:
	//	fmt.Println(data)
	//}

}
