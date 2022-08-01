package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//消费者消费
func consumer(queue chan int) {
	defer wg.Done()
	//data := <-queue
	for {
	    data,ok := <- queue
		if !ok{
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}
	//for data := range queue {
	//	fmt.Println(data)
	//	time.Sleep(time.Second)
	//}
	//fmt.Println(data)
}
func main() {
	/*
		channel 提供了一种通信机制 定向发消息 消息队列 python java
	*/
	//1. 定义一个channel
	var msg chan int
	//2. 初始化这个channel 两种方式
	msg = make(chan int) //第一种初始化方式 无缓冲
	//msg = make(chan int, 1)
	//第二种初始化方式 有缓冲空间
	//3. 在go语言中 使用make初始化的有三种 slice map channel
	//go func(queue chan int) {
	//	data := <-msg //将箭头右边的值放到左边
	//	fmt.Println(data)
	//}(msg)
	//msg <- 3 //fatal error: all goroutines are asleep - deadlock!
	wg.Add(1)
	go consumer(msg)
	msg <- 1 //你这个管道看起来好像就是一个有空间的数组
	//go consumer(msg)
	msg <- 2 //fatal error: all goroutines are asleep - deadlock!
	//关闭channel 1.已经关闭的channel不能再发送数据了 2.已经关闭的channel 消费者能继续取数据吗？
	//2可以 直到数据取完为止
	close(msg)
	//go consumer(msg)
	wg.Wait()

}
