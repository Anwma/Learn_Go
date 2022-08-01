package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(queue <-chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}
func stock(queue chan <- int)  {
	defer wg.Done()
	for  {
		queue <- 1
		time.Sleep(time.Second)
	}

}
func main() {
	//有没有缓冲 1.有缓冲 2.无缓冲
	//双向的还是单向的 为了安全 还提供了单向channel
	var msg chan int //双向 可以把双向的channel赋值给单向的channel
	//var msg chan<- int //仅发送
	//var msg <-chan int //仅接收
	msg = make(chan int, 10)
	//data := <-msg //无效运算: <-msg (从仅发送类型 chan<- int 接收)
	//fmt.Println(data)
	wg.Add(1)
	go consumer(msg) //普通的channel 可以直接转换为单向的channel
	msg <- 1
	close(msg)
	wg.Wait()
	/*
		1
		等待返回值
		2

		等待返回值
		1
		2
	*/
}
