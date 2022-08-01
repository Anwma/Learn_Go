package main

import "fmt"

func replace(mySlice []string) {
	mySlice[0] = "Wozen"
}
func main() {
	//什么是切片
	//数组有一个很大的问题 大小确定 不能修改 - 切片
	//var identifier []type
	//第一种：var courses []string //定义了一个切片
	//var courses = []string{"django", "scrapy", "tornado"}
	//fmt.Println(courses)
	//fmt.Printf("%T\n", courses)

	//切片的另一种初始化方法 make
	//切片不是没有长度限制 为什么使用make初始化的时候需要我们传递一个长度
	//那我传递了长度之后是不是意味着就像数组一样长度不能变了呢？
	//第二种: courses := make([]int, 10, 15)
	//println(cap(courses))
	//println(len(courses))
	//slice对标python中的list

	//第三种方法 通过数组 变成一个切片
	//var courses = [5]string{"django", "scrapy", "tornado", "python", "golang"}
	//subCourse := courses[1:4] //python中的用法叫切片 go语言中切片是一种数据结构
	//python中切片的用法非常灵活非常的多
	//切片
	//replace(subCourse)
	//fmt.Println(subCourse)
	//subCourse[0] = "hello"
	//fmt.Println(courses)
	//fmt.Printf("%T\n", subCourse)

	//第四种方式: new
	//subCourse := *new([]int) //[]
	//subCourse := new([]int) //&[]
	//fmt.Println(subCourse)

	//数组的传递是值传递 切片不是 切片是引用传递
	//slice很重要 很常用

	//slice是动态数组 所以说我们需要动态添加值
	//fmt.Println(subCourse[1])
	//subCourse[1] = "bobby"
	//fmt.Println(subCourse)
	//subCourse2 := subCourse[1:3]
	//fmt.Printf("%T, %v\n", subCourse2, subCourse2)

	//append 函数 可以向切片元素追加元素
	//subCourse2 = append(subCourse2, "anwma1", "anwma2", "anwma3")
	//fmt.Println(subCourse2)

	//append函数追加多个元素
	//appendCourse := []string{"anwma1", "anwma2", "anwma3"}
	//subCourse2 = append(subCourse2, appendCourse...)
	//subCourse3 := make([]string, 2)
	//subCourse3 := make([]string, len(subCourse2))
	//fmt.Println(len(subCourse3))
	//copy(subCourse3, subCourse2)

	//拷贝的时候 目标对象长度需要设置好
	//fmt.Println(subCourse3)

	//append 函数 既然是切片 那为什么这个时候又来提出长度的要求
	//想从slice中删除元素怎么办
	//deleteCourses := [5]string{"django", "scrapy", "tornado", "python", "golang"}
	//courseSlice := deleteCourses[:]
	//courseSlice = append(courseSlice[:1], courseSlice[2:]...) //取巧的做法
	//fmt.Println(courseSlice)

	//如何判断某个元素在切片中

	//python和go的slice区别 玩出花样 go的slice更像是python的list
	//go语言的slice 是基于数组实现的 python的list底层也是基于数组实现的(C)
	//slice进行的操作都会影响原来的数组 slice更像是一个指针 本身不存值

	//slice的原理 因为很多底层的知识相对来说很多时候并不难而是需要花费比较多的时间去慢慢理解
	//1.第一个现象
	a := make([]int, 0)
	b := []int{1, 2, 3}
	fmt.Println(copy(a, b)) //0
	fmt.Println(a)          //[]
	//2.第二个现象 修改切片的时候影响原数组
	c := b[:]
	//c[0] = 8
	fmt.Println(b) //[1 2 3]
	fmt.Println(c) //[1 2 3]
	//3.第三个现象
	c = append(c, 9)
	fmt.Println(b) //[1 2 3] //append函数没有影响到原来的数组
	fmt.Println(c) //[1 2 3 9]
	c[0] = 8
	fmt.Println(b) //[1 2 3]
	fmt.Println(c) //[8 2 3 9]//为什么append函数之后再调用c[0] = 8不会影响原来的数组呢
	//4.第四个现象
	fmt.Println(len(c))
	fmt.Println(cap(c)) //cap指的是容量 长度 容量这是两个不同的概念
	//切片底层是使用数组实现的 既要使用数组 又要满足动态的功能 怎么实现？
	//假设有一个值 实际上申请数组的时候可能是两个
	//如果后续要增加数据那么就直接添加到数据的结尾	这个时候我不要额外重新申请
	//切片有不同的初始化方式
	//1.使用make方法初始化 len 和 cap是多少
	d := make([]int, 5, 6)
	fmt.Printf("len=%d,cap=%d\n", len(d), cap(d))

	//2.通过数组取切片
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[2:4]
	newSlice := data[3:6]
	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Printf("len=%d,cap=%d\n", len(slice), cap(slice))
	fmt.Printf("len=%d,cap=%d\n", len(newSlice), cap(newSlice))

	//3.
	slice2 := []int{1, 2, 3}
	fmt.Printf("len=%d,cap=%d\n", len(slice2), cap(slice2))

	//切片扩容问题 扩容阶段会影响速度 python中的list中底层实际上也是数组 也会面临动态扩容的问题
	//python的list中数据类型可以不一致
	oldSlice := make([]int, 0)
	fmt.Printf("len=%d,cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 0)
	fmt.Printf("len=%d,cap=%d\n", len(oldSlice), cap(oldSlice))
	/*
		Go中切片的扩容机制

		首先判断，如果新申请（cap）大于2倍的就容量，最终容量（newcap)就是新申请的容量
		否则判断，如果旧切片的长度小于1024 则最终容量（newcap)就是旧容量（old.cap）的两倍，即（newcap=doublecap）
		否则判断，如果旧切片长度大于等于1024 则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4
		即（newcap）
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
	*/
	//如果小于1024 扩容速度是两倍 如果大于1024 扩容速度就是1.25倍

	//切片来说
	//1.底层是数组 如果是基于数组产生的 会有一个问题就是会影响原来的数组
	//2.切片的扩容机制
	//3.切片的传递是引用传递
	oldArr := [3]int{1, 2, 3}
	newArr := oldArr
	newArr[0] = 5
	fmt.Println(newArr, oldArr)

	//go语言中slice的原理讲解很重要 有坑（对于初学者来说）
	//有经验的程序觉得这个不是坑
	//程序员也是消费者 - java c++ go python -才会站在使用者的角度上去考虑 对于处理器就不友好
	//当make遇上了append容易出现的坑
	s1 := make([]int, 5)

	s1 = append(s1, 6)
	fmt.Println(s1)
	//很多人对make函数产生了一个假象 s1 := make([]int, 5) 好比是python的s1 = []
}
