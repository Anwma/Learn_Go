package main

import (
	"fmt"
)

func main() {
	//f, _ := os.OpenFile("xxx.txt")
	//此处有大量的逻辑需要读取文件
	//defer f.Close()
	/*
		defer 语句是go提供的一种用于注册延迟调用的机制 它可以让当前函数执行完毕之后执行
		对于python的with语句来说
	*/
	//fmt.Println("hello1")
	//defer a++ error 可以做函数调用不能做表达式的书写 比如a++
	//a := 0
	//defer fmt.Println("defer hello")
	//panic("error")
	//fmt.Println("hello2")

	//如果有多个defer会出现什么情况
	//defer fmt.Println("test1")
	//defer fmt.Println("test2")
	//defer fmt.Println("test3")
	//多个defer是按照先入后出的顺序执行的

	//defer 语句执行时的拷贝机制
	//test := func() {
	//	fmt.Println("test1")
	//}
	//defer test()
	//test = func() {
	//	fmt.Println("test2")
	//}
	//fmt.Println("test3")
	//
	//x := 10 //压栈的是函数段
	//defer func(a int) {
	//	fmt.Println(x) //11
	//	fmt.Println(a) //10
	//}(x)
	//x++
	//defer func(a *int) {
	//	fmt.Println(*a)
	//}(&x)
	//x++
	//闭包 此处的defer函数没有参数 - 函数内部使用的值是全局的值
	//defer func() {
	//	fmt.Println(x) //11
	//}()
	//x++
	//defer func(x int) {
	//	fmt.Println(x) //10
	//}(x)
	//x++
	fmt.Println(f1()) //10 是不是就意味着 defer中影响不到外部的值呢？ - 不是
	fmt.Println(*f2())
	//defer本质上是注册了一个延迟函数 defer函数的执行顺序已经确定
	//defer没有嵌套 defer的机制是要取代try except finally的机制
}
func f1() int {
	x := 10
	defer func() {
		x++ //闭包里面的值 不影响
	}()
	tmp := x
	return tmp
	//return x
}
func f2() *int {
	a := 10
	b := &a
	defer func() {
		*b++ //闭包里面的值 不影响
	}()
	temp_data := b
	return temp_data
	return b
}
