package main

//采用字符串
import (
	//int_calc "PackageTest/calc"
	_ "PackageTest/calc"
	//float_calc "PackageTest/calc_float"
	//"fmt"
)

//1.新建项目的时候 选择go modules
//2.新建一个包的时候 新建一个文件夹-一般小写 这个文件下下面的所有的源码中的package都一样
//3.import "PackageTest/calc" 这个实际上很多人会认为就是包名 但是实际上是包的路径
//4.go同一个文件夹下面的package名称可以不和文件夹名称保持一致 但是这些源码中的package名称必须一样
//5.在python中可以import一个包中的函数 变量
func main() {
	//fmt.Println(int_calc.Add(1, 2))
	//fmt.Println(float_calc.Add(1.21, 2.45))
	//fmt.Println("hello go!")
	//go的包管理和python的包管理看起来差不多
}
