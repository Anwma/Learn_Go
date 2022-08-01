package main

import "fmt"

func swap(a *int, b *int) {
	//用于交换a和b
	c := *a
	*a = *b
	*b = c
}
func main() {
	//什么是指针 我们提一个问题
	a := 10
	b := 20
	//swap(a, b)
	fmt.Println(a, b)
	//为什么交换不成功 这个函数运行完成以后 我想要把a 和 b的值变掉
	//指针 -对于内存来说 每一个字节其实都是有地址的 - 通过16进制打印出来
	fmt.Printf("%p\n", &a) //变量有地址
	//现在有一种特殊的变量类型 这个变量只能保存地址值
	var ip *int //这个变量里面只能保存地址类型这种值
	ip = &a

	//如果要修改指针指向的变量的值, 用法也比较特殊
	*ip = 30
	fmt.Println(a)
	//如果定义变量 如何修改指针变量指向的内存中的值 通过指针去取值的时候不知道应该取多大的连续内存空间
	fmt.Printf("ip所指向的内存空间地址是: %p, 内存中的值是: %d\n", ip, *ip)
	swap(&a, &b)
	fmt.Println(a, b)
	//还不足以说服大家 但是go中数组是值传递
	//数组中有100万个值 对于这种一般我们都采用切片来传递
	//在python中list和dict这种传递都是引用传递

	//指针还可以指向数组 指向数组的指针 数组是值类型
	//arr := [3]int{1, 2, 3}
	//var ip_arr *[3]int = &arr
	//指针数组
	//var ptrs [3]*int //创建能够存放三个指针变量的"数组"
	//很多时候都是函数参数指明的类型
	//指针的默认值是nil
	if ip != nil {

	}
	//if a != nil {
	//
	//}

	//像python 和 java这种语言都在极力的屏蔽指针
	//C/C++ 都提供了指针 指针本身是很强大的
	//C/C++中指针的功能很强大 指针的转换/偏移/运算
	//go语言没有屏蔽指针,但是go语言在指针上做了大量的限制,安全性高很多,相比较 C/C++ 灵活性就降低了
	//go的指针变量中涉及到两个符号 & 和 *

}
