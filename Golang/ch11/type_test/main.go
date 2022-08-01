package main

import "fmt"

func main() {
	//go语言中的关键词 type
	//1.给一个类型定义别名,实际上为什么会有byte 就是我为了强调我们现在处理的对象是字节类型
	//2.这种别名实际上还是为了代码的可读性,这个实质上本质仍然是uint8 无非就是在代码编码阶段可读性强而已
	//rune  byte
	type myByte = byte
	var b myByte
	fmt.Printf("%T\n", b)

	//3.第二种就是基于一个已有的类型定义一个新的类型
	type myInt int
	var i myInt
	fmt.Printf("%T\n", i)

	//4.定义结构体
	type Course struct {
		name  string
		price int
	}
	//5.定义接口、
	type Callable interface {
	}
	//6.定义函数别名
	type handle func(str string)
}
