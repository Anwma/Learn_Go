package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
使⽤两个  goroutine 交替打印序列，⼀个  goroutine 打印数字， 另外⼀
个  goroutine 打印字⺟, 最终效果如下：

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
*/
func main() {
	number, letter := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			<-number
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- true
		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-letter
			if i >= strings.Count(str, "")-1 {
				wait.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			number <- true
		}
	}(&wait)
	number <- true
	wait.Wait()
}
