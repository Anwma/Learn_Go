package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1. 基本的类型转换
	//a := int(3.0)
	//fmt.Println(a)
	////在go语言中不支持变量间的隐式类型转换
	////1.变量间类型转换不支持
	////var b int = 5.1 //5.0是常量 常量与变量支持 要求严格
	////fmt.Println(b)
	//c := 5.0
	//fmt.Printf("%T\n", c)
	////两种类型不一致
	////var d int = c
	//var d = int(c)
	//fmt.Println(d)

	//var a int64 = 56
	//var b = int32(a)
	//fmt.Println(b)

	//int转字符串 itoa
	//fmt.Printf("%T\n", strconv.Itoa(int(a)))
	//字符串转int atoi
	//data, _ := strconv.Atoi("12")
	//fmt.Println(data)
	//parse类的函数 字符串转bool float int Uint
	//b, err := strconv.ParseBool("True")
	//fmt.Println(b, err)
	//f, err := strconv.ParseFloat("3.1415", 32)
	//fmt.Println(f, err)
	//fmt.Printf("%T\n", f)
	//base 将字符串中的数字进行解析, 以base的进制进行解析
	//i, err := strconv.ParseInt("-12", 8, 64)
	//fmt.Println(i, err)
	//u, err := strconv.ParseUint("42", 10, 64)
	//fmt.Println(u, err)

	//其他类型转字符串
	s1 := strconv.FormatBool(true)
	fmt.Println(s1)
	//true
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Println(s2)
	//3.1415E+00
	s3 := strconv.FormatInt(-42, 16)
	fmt.Println(s3)
	//表示将-42转换为16进制数，转换的结果为-2a。
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s4)
	//2a
	var data int
	var err error
	if data, err = strconv.Atoi("12"); err != nil {
		fmt.Println("转换出错")
	}
	fmt.Println(data)
}
