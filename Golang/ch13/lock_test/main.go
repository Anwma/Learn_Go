package main

import (
	"fmt"
	"sync"
)

/*
	锁 - 资源竞争
	1.按理说 最后的结果是0
	2.实际的情况 1. 不是0 2.每次运行的结果不一样
*/
var total int
var wg sync.WaitGroup
//互斥锁(同步问题) - 读写锁 同步数据 能不用锁就别用锁 - 性能下降
//绝大多数的web系统来说 都是读多写少
//有1w个人同时读数据库 A读的时候 B能读么？ 为什么要加锁呢？100 - 1w - 100
//立马我下单了 - 支付（会重新查一遍数据库 1w）
//一定要加锁 写 和 读上面要加同一把锁 并发严重下降。
//B读了一个数据,造成C读了数据产生影响吗? 一定是写和读之间造成的
//如果这把锁之间可以做到 读之间不会产生影响 读写之间才会产生影响 那多好 读写锁

var lock sync.Mutex
func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//先把门锁上
		lock.Lock()
		total += 1
		//放开锁
		lock.Unlock()
		//1.从total中取值
		//2.将total+1
		//3.将total+1的计算结果放入到total中
	}
}
func sub(){
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//先把门锁上
		lock.Lock()
		total -= 1
		lock.Unlock()
		//放开锁
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
