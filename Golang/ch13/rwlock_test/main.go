package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	锁 - 资源竞争
	1.按理说 最后的结果是0
	2.实际的情况 1. 不是0 2.每次运行的结果不一样
*/
var total int
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁
var rwLock sync.RWMutex //读写锁
func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取成功")
	rwLock.RUnlock()

}
func write(){
	defer wg.Done()
	rwLock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second*2)
	fmt.Println("修改成功")
	rwLock.Unlock()

}
func main() {
	wg.Add(8)
	for i := 0; i < 5; i++ {
		go read()
	}
	for i := 0; i < 3; i++ {
		go write()
	}
	wg.Wait()
	fmt.Println(total)
}
