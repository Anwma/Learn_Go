package main

import (
	"fmt"
	"time"
)

func f1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获到了")
		}
	}()
	go func() {
		panic("出错了")
	}()
	time.Sleep(2 * time.Second) //为啥panic了？
}
func main() {
	f1()
	//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
	//a := 10
	//b := 0
	//fmt.Println(a / b)
	//panic("error")
	//go func() {
	//操作redis的，有人觉得这段代码可以放到协程中去运行。有一个非常大的隐患了
	//panic("error")
	//}()
	//服务被停掉
	//writer.Write([]byte("hello world"))
	//})
	//http.ListenAndServe("127.0.0.1:8080", nil)
	//panic会引起主线程的挂掉 同时会导致其他的协程都挂掉
	//为什么直接在,在父协程中无法捕获走子协程出现的异常

}
