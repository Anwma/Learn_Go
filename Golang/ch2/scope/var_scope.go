package main

import "fmt"

var c = 20

func main() {
	//变量的作用域
	//局部变量
	//缩进的习惯
	//静态语言写起来代码多,但是严谨性很好
	//c := 10 //全局变量定义时无法使用短变量声明
	//fmt.Println(c)
	////Output: 10
	//sex := "Female"
	//if sex == "Female" {
	//	outStr := "女"
	//}
	//fmt.Println(outStr)
	//程序无法运行
	fmt.Printf("%c\n", 97)
	//在go中字符和字符串不是一种类型 字符类型是单引号 字符串是双引号
	fmt.Printf("%d", '慕') //int32 Unicode编码
}
