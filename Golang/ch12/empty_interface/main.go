package main

import "fmt"

//type Empty interface {
//}

type Course struct {
	name  string
	price int
	url   string
}
type Printer interface {
	printInfo() string
}

func (c Course) printInfo() string {
	return "课程信息"
}

//func print(x interface{}) {
//	//此处的ok是bool类型
//	if v, ok := x.(int); ok { //判断是不是int类型
//		//x是int类型
//		fmt.Printf("%d, 整数\n", v)
//	}
//
//	if s, ok := x.(string); ok {
//		fmt.Printf("%s,字符串\n", s)
//
//	}
//	//牵扯到go的另一个默认的问题
//	//fmt.Printf("%v\n	", i)
//}

type AliOss struct {
	//保存
	//下载
}
type LocalFile struct {
}

func store(x interface{}) {
	switch v := x.(type) {
	case AliOss:
		//此处要做一些特殊的处理,比如设置阿里云的权限问题
		fmt.Println(v)
	case LocalFile:
		//检查路径的权限
		fmt.Println(v)
	}

}

func print(x interface{}) {
	//类型断言
	switch x := x.(type) {
	case string:
		fmt.Printf("%s,字符串\n", x)
	case int:
		fmt.Printf("%d, 整数\n", x)
	case float64:
		fmt.Printf("%f, 浮点数\n", x)
	}
}

func main() {
	//空接口
	var i interface{} //有点像多态
	//空接口可以类似于我们java和python中的object
	//i = Course{}
	//fmt.Println(i)
	i = 10
	print(i)
	i = "hello"
	fmt.Println(i)
	//i = []string{"django", "scrapy"}
	//fmt.Println(i)
	//1.空接口的第一个用途,可以把任何类型都赋值给空接口变量

	//2.参数传递 什么类型都可以打印 什么类型都可以接收

	//3.空接口可以作为map的值
	//var teacherInfo = make(map[string]interface{})
	//teacherInfo["name"] = "wozen"
	//teacherInfo["age"] = 18
	//teacherInfo["weight"] = 72.5
	//teacherInfo["courses"] = []string{"hello", "world"}
	//fmt.Printf("%v\n", teacherInfo)
	//类型断言
	//接口的一个坑, 接口引入了
	//接口还有一个默认的规范 接口的名称以 er 结尾
	//c := &Course{}
	//var c Printer = Course{}
	//c.printInfo()
	//fmt.Println()

}
