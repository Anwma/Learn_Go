package main

import (
	"fmt"
	"time"
)

//go的协程和python的协程 - 网上有些人可能喜欢拿web框架来做性能对比
//go的gin beego - flask/django + gevent性能对比 - 这个不科学的 uwsgi部署
//flask/django 只是一个web框架 本身不提供 启动多线程 多进程来提高并发的能力 一般使用中间件
//tornado sanic fastapi /asyncio 协程库 我们就不再使用python的多线程来对比了
//只要大家懂了任何一门语言的协程 其他语言的协程都很好理解 GMP调度模型
//1.大家都开启100w个协程 2.使用的简单性
func p() {
	//fmt.Println("协程")
	for i := 1; i <= 100000; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
}
func main() {
	//闭包的特性
	//主死从随
	//go p()
	for i := 1; i <= 300000; i++ {
		go func(n int) {
			fmt.Println(n)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second*10)
}

//goroutine
//python java c++ 多线程和多线程编程
//go 语言诞生的比较晚 web 2.0 开发主键主流 高并发
//多线程 - 每个线程占用的内存比较多 而且系统切换开销很大 上千绿程/轻量级线程 - 协程 - 用户态线程
//go语言一开始的时候就没有让我们取实例化一个线程 - 协程 go 没有历史包袱
//nodejs python3.6 那为什么go的协程为什么会这么火
//python中有两种编程模式 1.多线程和多进程来进行并发编程 2. 使用协程进程并发编程 很多的库是基于多线程和多进程开发的
//除非某一天所有的库 大部分的库都支持了协程
