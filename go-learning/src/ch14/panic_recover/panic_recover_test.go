package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("Finally!")
	//}()

	/*
		其实，这样的修复方式是非常危险的。我们一定要当心我们自己的recover在做什么事，
		因为recover的时候，并不去检测错误到底发生了什么错误。而是我们直接把错误，记录到日志文件当中
		或者说是忽略掉了。这时候可能是系统里面的某些核心资源消耗完了，那么我们把他强制的恢复掉之后呢，
		其实系统依然是不能够正常工作的，他会导致我们的一些健康检查程序没有办法检查出当前系统的问题。
		因为很多的health check 他只是检查当前系统进程在还是不在，因为我们的进程是在的，所以说就会导致
		一种僵尸服务进程。它好像活着，但其实它也不能提供服务，这种情况下呢，我个人认为你倒不如采用一种
		可恢复的设计模式。其中一种叫做“Let's it crash”。我们干脆让他 crash 掉。一旦crash掉，我们的
		守护进程，就会帮我们重新把我们的服务进程提起来。重启有时候就像以前我们windows不稳定的时候。
		重启有时候是一种恢复不确定性最好的方法。
	*/
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
	//os.Exit(-1)
	//fmt.Println("End")
}
