package main

import "fmt"

//省略号 ...
func add(params ...int) (sum int) {
	//不能解决一个问题我可能有不定个int值传递进来
	for _, v := range params {
		sum += v
	}
	return
}
func add2(params []int) (sum int) {
	//不能解决一个问题我可能有不定个int值传递进来
	for _, v := range params {
		sum += v
	}
	params[0] = 9
	return
}

type sub func(a, b int) int //sub就等同于 int map
func subImpl(a, b int) int {
	return a - b
}
func filter(score []int, f func(int) bool) []int {
	reSlice := make([]int, 0)
	for _, v := range score {
		if f(v) {
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}
func main() {
	//通过省略号去动态设置多个参数
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(add(1, 2, 3, 4, 5))
	//这种效果 - slice
	//区别 slice是一种类型 还是引用传递
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(add2(slice))
	fmt.Println(add(slice...)) //将slice打散
	fmt.Println(slice)
	//我们要慎重

	//省略号的用途
	//1. 函数参数不定长 2.将slice打散 3.
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr)
	//函数内部定义函数 函数内嵌
	//匿名函数
	fmt.Println(func(a, b int) int {
		return a + b
	}(1, 2))

	//fmt.Printf("%T\n", result) //func(...int) int

	//go语言中非常重要的特性 函数 一些trick 一些语法糖
	//函数的一等公民特性 - 可以作为参数 返回值 复制给另一个变量
	//函数也可以当作参数传递给一个函数
	var mySub sub = func(a, b int) int {
		return a - b
	}
	fmt.Println(mySub(1, 2))
	//mySub(1, 2)

	//将函数作为另一个函数的参数
	//写一个函数用于过滤一部分数据
	score := []int{10, 30, 40, 50, 90}
	//写一个函数 过滤掉不合格的成绩
	fmt.Println(filter(score, func(a int) bool {
		if a >= 90 {
			return true
		} else {
			return false
		}
	}))
	//gin python的装饰器

	//go语言并没有提供try except finally
	//1.大量嵌套 try finally 2.打开和关闭逻辑离得太远
}
