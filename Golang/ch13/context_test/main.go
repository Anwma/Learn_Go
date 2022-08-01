package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1.监控全局变量 来完成
//2.通过channel
//var stop bool
//var stop chan bool = make(chan bool)
//3.刚才的两种方式 这种方案并不统一 go1.7 context机制
//父的context
//父的context取消 那么这个父的context生成的context也会被取消
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	//ctx2, _ := context.WithCancel(ctx)
	context.WithDeadline(ctx, time.Now())
	//a -> b -> c
	//context.WithTimeout()
	//context.WithValue()

	//go memoryInfo(ctx2)
	for {
		//if stop {
		//	break
		//}
		//select {
		//case <-stop:
		//	fmt.Println("退出CPU监控")
		//	return
		//default:
		//	time.Sleep(time.Second*2)
		//	fmt.Println("CPU信息读取完成")
		//}
		select {
		case <-ctx.Done():
			fmt.Println("监控退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取cpu信息成功")
		}
	}
}

func memoryInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("监控内存退出")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("获取内存信息成功")
		}
	}
}
func main() {
	//现在启动一个goroutine去监控某台服务器的cpu使用情况
	wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2) //有一种自动的机制在里面
	//ctx, cancel := context.WithDeadline(context.Background(), time.Second*2) //有一种自动的机制在里面
	go cpuInfo(ctx)
	time.Sleep(time.Second)
	cancel()
	//go memoryInfo(ctx)
	//现在希望可以中断CPU的信息读取
	//time.Sleep(time.Second * 5)
	//cancel()
	//context在web开发中非常的常用 grpc 每个接口调用都会传递context gin的http接口也会有
	//stop <- true
	//stop = true
	wg.Wait()
	fmt.Println("信息监控完成")
}
