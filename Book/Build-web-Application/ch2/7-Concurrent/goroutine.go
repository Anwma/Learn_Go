package main

import (
	"fmt"
	"runtime"
)

//有一种情况是这样的，正常执行 Go 程序就会出现问题，而断点debug跟踪此方法的时候却是正常的，
//不断测试结果都是这样，由此判断有可能是因为此方法在执行的时候所需要的参数在获取的时候需要一段时间，
//而debug的时候是一步一步执行代码，时间很充足，而当程序正常执行的时候由于执行的时间很快，
//某一个或一些参数没有获取到就直接执行了此方法，由此导致执行结果就会出现和debug时候的执行结果不一致的问题，
//针对这样的问题有一个解决方法，就是在执行这个方法之前让程序停一会儿，给获取需要的参数所执行的代码足够的时间。
//给调用此方法的代码之前加上下面语句 time.Sleep()

//建议加锁??? 结果是不确定的

//这个函数的作用是在自己控制的cpu范围之内让出cpu执行权限，而不是无限制让出；

func say(s string) {

	for i := 0; i < 3; i++ {
		//Gosched 让出处理器，允许运行其他goroutines。它不会挂起当前的goroutine，因此会自动继续执行。
		runtime.GOMAXPROCS(800)
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
	//time.Sleep(time.Second)
}

