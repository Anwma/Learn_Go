package main

import "fmt"

func main() {
	//定义bool类型
	//var gender = false
	//fmt.Println(gender)
	//相比python而言，go语言为什么有这么多种整数类型
	//年龄、分数都是有上限的年龄不超过0-200 分数不超过0-150
	//很多场景之下，数字有上限，我们可以选择合适的数据类型来降低内存的占用
	//int是一种动态类型取决于机器本身
	//var age int16 = 18
	//fmt.Println(unsafe.Sizeof(age))
	//fmt.Println(age)

	//float类型
	//var weight float32 = 71.2
	//fmt.Println(weight)
	//fmt.Println(math.MaxFloat32)
	//fmt.Println(math.MaxFloat64)
	//为什么64位的float最大值远大于int64位，float底层存储和int的存储不一样
	//float32 float64 两者占用的内存不一样 64位的最大数和精度都比32位高
	//weight := 71.2
	//fmt.Printf("%T\n", weight)
	//age := 18
	//fmt.Printf("%T\n", age)

	//byte 和 rune
	//byte 类型
	//静态语言中 中文处理很容易出错
	//var a byte = 18
	//fmt.Println(a)

	a := 'a'
	//这里注意一下
	//1. a+1可以和数字计算 2.a+1的类型是32 3. int类型可以直接变成字符
	fmt.Printf("%T\n", a+1)
	fmt.Printf("a+1 = %c", a+1)
}
