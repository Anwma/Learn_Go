package main

import "fmt"

func test() (int, error) {
	return 0, nil
}

func main() {
	//变量的定义
	//静态语言的变量和动态语言的变量差异很大
	//1. 变量的声明和定义
	//var a = 20
	//var b = 20
	//c := 100
	//fmt.Println(a, b, c)
	//
	//d, e := 1, 2
	//fmt.Println(d, e)
	//
	//var f, g, h int
	//f, g, h = 1, 2, 3
	//fmt.Println(f, g, h)
	////集合类型
	//var (
	//	a    int
	//	name string
	//)

	var i = 10
	i = 20
	fmt.Println(i)

	//匿名变量 变量一旦被定义 不使用会报错
	//常量的定义
	//const PI = 3.14159265
	//r := 2.0
	//c := 2. * PI * r
	//fmt.Println(c)

	//枚举
	//const (
	//	Unknown = 0
	//	Female  = 1
	//	Male    = 2
	//)

	//常量组如不指定类型和初始化值,该类型和值和上一行的一致
	//const (
	//	x int = 16
	//	y
	//	s = "abc"
	//	z
	//)
	//fmt.Println(x, y, s, z)
	//Output:16 16 abc abc
	//1. 常量的数据类型值可以是布尔、数字、字符串
	//2. 不曾使用的常量，在编译的时候不会报错
	//const常量的iota, 常量计数器
	const (
		Unknown = iota //计数器 从0开始 第一行不能省略
		Female
		Male
		//Book    = 0
		//Cloth   = 1
		//Phone   = 2
		//DestTop = 3
	)
	//0,1,2 本身不重要，重点在于这三个值不一样
	fmt.Println(Unknown, Female, Male)
	//iota你真的懂了吗?
	//	1. iota只能在常量组中使用
	//	2. 各个常量组之间,const定义块互不干扰
	// 	3. 没有表达式的常量定义复用上一行的表达式
	//	4. 从第一行开始,iota从0 "逐行" 加一
	const (
		a = iota //iota = 0
		b = 10   //iota = 1 每一行iota加一
		c        // c = 10, iota = 2
		//d = iota
		d, e = iota, iota
		f    = iota
		//iota代表的是这里的行数
	)
	//fmt.Println(a, b, c, d)
	//Output: 0 10 10 3
	fmt.Println(a, b, c, d, e, f)
	//Output: 0 10 10 3 3 4
}
