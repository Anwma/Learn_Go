package main

import "fmt"

//匿名函数调用
func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		// 将 匿名函数func(i int) 赋值给 g
		g := func(i int) { fmt.Printf("%d ", i) }
		//调用g 匿名函数  输出g 中的内容 fmt.Printf("%d ", i)
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}
