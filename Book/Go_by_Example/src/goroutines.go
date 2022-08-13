package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		//解锁fmt.Println用法
		//func Println(a ...any) (n int, err error) {
		//		return Fprintln(os.Stdout, a...)
		//}
		fmt.Println(from, ":", i)
	}
}

func main() {
	//并发执行direct 和 goroutine
	//异步进行
	f("direct")
	//go f(s) 在一个协程中调用这个函数。 这个新的 Go 协程将会 并发地 执行这个函数。
	go f("goroutine")
	//going的位置会不确定
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
