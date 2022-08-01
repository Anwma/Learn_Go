package main

import (
	"fmt"
	"sync"
)
//如何解决主的goroutine在子协程结束后自动结束
var wg sync.WaitGroup
//WaitGroup 提供了三个很有用的函数
/*
	Add
	Done
	Wait
	Add的数量必须和Done的数量相等
 */
func f(i int)  {
	//如果不写，fatal error: all goroutines are asleep - deadlock!
	//defer wg.Done()
	fmt.Println(i)
}
func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go f(i)
		//go func(n int) {
			//一般利用defer机制
			//defer wg.Done()
			//fmt.Println(n)
			//wg.Done()
		//}(i)
	}
	wg.Wait()
}
