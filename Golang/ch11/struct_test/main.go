package main

import (
	"fmt"
	"unsafe"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

// 结构体绑定方法 函数的接收者
func (c Course) printCourseInfo() {
	fmt.Printf("课程名: %s, 课程价格: %d, 课程的地址: %s", c.Name, c.Price, c.Url)
}
func (c *Course) setPrice(price int) {
	c.Price = price
}

//结构体的方法只能和结构体在同一个包中
//2.内置的int类型不能加方法

func main() {
	//go语言不支持面向对象
	//面向对象的三个基本特征：封装 继承 多态      方法重载 抽象基类(不支持)
	//定义struct go语言没有class这个概念 所以说对于很多人来说会少理解很多面向对象抽象的概念

	//1.实例化 k-v方式
	//var c new_pkg.Course = new_pkg.Course{
	//	Name:  "django",
	//	Price: 100,
	//	Url:   "https://www.imooc.com",
	//}
	//访问
	//fmt.Println(c.Name, c.Price, c.Url)
	//大小写在go语言中的重要性 可见性
	//一个包中的变量或者结构体如果首字母是小写，那么对于另一个包不可见
	//机构体定义的 名称 以及属性首字母大写很重要

	//2.第二种实例化方式 - 顺序形式
	//c2 := Course{"scrapy", 110, "https://www.imooc.com"}
	//fmt.Println(c2.Name, c2.Price, c2.Url)
	//3.如果一个指向结构体的指针 通过结构体指针获取对象的值
	//c3 := &Course{"tornado", 100, "https://www.imooc.com"}
	//fmt.Printf("%T\n", c3)
	//fmt.Println((*c3).Name, (*c3).Price, (*c3).Url)
	//fmt.Println("--------------")
	//应该能看出来 go语言实际上在借鉴很多动态语言的特性 -很多地方不管如何写都是正确的
	//另一个根本原因 go语言的指针是受限的
	//fmt.Println(c3.Name, c3.Price, c3.Url) //这里其实是go语言的一个语法糖 go语言内部会将c3转换为(*c3)
	//4.零值 如果不给结构体赋值 go语言会默认给每个字段采用默认值
	//c4 := Course{}
	//fmt.Println(c4.Name)

	//5.多种方式零值初始结构体
	//var c5 Course = Course{}
	//var c6 Course
	//var c7 *Course = new(Course)
	//var c8 *Course
	//var c9 *Course = &Course{}
	//为什么c6和c8表现出来的结果不一样 指针如果只声明不赋值 默认值是nil c6不是指针 是结构体的类型
	//fmt.Println("零值初始化")
	//fmt.Println(c5.Price)
	//fmt.Println(c6.Price)
	//fmt.Println(c7.Price)
	//fmt.Println(c8.Price)
	//fmt.Println(c9.Price)

	//6.结构体是值类型
	c10 := Course{"scrapy", 110, "https://www.imooc.com"}
	c11 := c10
	fmt.Println(c10)
	fmt.Println(c11)
	c10.Price = 200
	fmt.Println(c10)
	fmt.Println(c11)
	//go语言中struct无处不在
	//7.结构体的大小 占用内存的大小 可以使用sizeof来查看对象占用的类型
	fmt.Println(unsafe.Sizeof(1)) //8
	//go语言string的本质 其实string是一个结构体 16个字节
	//type string struct {
	//	Data uintptr //指针占8个长度
	//	Len  int     //长度64位系统占8个长度
	//}
	fmt.Println(unsafe.Sizeof("scrapy")) //16
	fmt.Println(unsafe.Sizeof(c10))      //16

	//8.slice的大小 24个字节
	type slice struct {
		array unsafe.Pointer //底层数组地址
		len   int            //长度
		cap   int            //容量
	}
	s1 := []int{1, 2, 4, 5}
	s2 := []string{"hello"}
	fmt.Println("切片的大小:")
	fmt.Println(unsafe.Sizeof(s1))
	fmt.Println(unsafe.Sizeof(s2))

	m1 := map[string]string{
		"hello1": "world1",
		"hello2": "world2",
		"hello3": "world3",
	}
	fmt.Println(unsafe.Sizeof(m1)) //8
	//结构体方法 达到了封装数据和封装方法的效果
	c10 = Course{"scrapy", 100, "https://www.imooc.com"}
	//Course.printCourseInfo(c10)
	//Course.setPrice(c10, 200)
	(&c10).setPrice(200) //修改c10的price？为什么呢？ 内部转换 语法糖 函数参数的传递是怎么传递的 结构体是值传递
	fmt.Println(c10.Price)
	//c10.printCourseInfo()
	//结构体的接收者有两种形式 1.值传递2.指针传递 如果你想改结构体的值 如果结构体的数据很大
	//go语言不支持继承 但是有办法能达到同样的效果 组合

}
