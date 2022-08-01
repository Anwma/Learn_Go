package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//Gosched 让出处理器，允许运行其他goroutines。它不会挂起当前的goroutine，因此会自动继续执行。
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
}
