package main

import "fmt"

func main() {

	//通道(channels) 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。
	//使用 make(chan val-type) 创建一个新的通道
	messages := make(chan string)
	//channel(messages) <- string 发送一个值到通道中
	go func() { messages <- "ping" }()
	//从通道中 接收一个值
	//变量 := <-channel
	msg := <-messages
	fmt.Println(msg)
}
