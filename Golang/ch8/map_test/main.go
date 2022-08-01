package main

import "fmt"

func main() {
	//go语言中的map ->python中的dict
	//go语言中的map的key和value类型声明的时候就要指明
	//1.字面值
	m1 := map[string]int{
		"m1": 1,
	}
	fmt.Printf("%v\n", m1)

	//2.make函数 make函数可以创建slice 可以创建map
	m2 := make(map[string]int) //创建时不添加元素
	m2["m2"] = 2
	fmt.Printf("%v\n", m2)

	//3.定义一个空的map
	m3 := map[string]string{}
	fmt.Printf("%v\n", m3)

	//map中的key不是所有的类型都支持，该类型需要支持 ==或者 != 操作
	//int rune
	//a := []int{1, 2, 3}
	//b := []int{1, 2, 3}
	//两个切片之间不能比较
	//if a == b {
	//
	//}
	//var m1 map[[]int]string

	//数组支持比较
	//a := [3]int{1, 2, 3}
	//b := [3]int{1, 2, 3}
	//if a == b {
	//	println("yes")
	//}

	//map的基本操作
	m := map[string]string{
		"a": "va",
		"b": "vb",
		"d": "",
	}

	//1.进行增加 修改
	m["c"] = "vc"
	m["b"] = "vb1"
	fmt.Printf("%v\n", m)

	//查询 你返回空的字符串到底是没有获取到还是值本身就是[]
	v, ok := m["d"]
	if ok {
		fmt.Println("找到了", v)
	} else {
		fmt.Println("没找到", v)
	}
	fmt.Println(v, ok)

	//删除
	//delete(m, "a")
	//delete(m, "e") //删除不存在的元素时 不抛异常
	//delete(m, "a") //删除已经删除过的元素时 不抛异常
	//fmt.Printf("%v\n", m)

	//遍历
	for k, v := range m {
		fmt.Println(k, v)
	}

	//go语言中也有一个list 就是数据结构中提到的链表
	//指针 为什么指针在java python等很多语言中不存在

}
