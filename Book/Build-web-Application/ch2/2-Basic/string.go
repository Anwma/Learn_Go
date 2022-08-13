package main

import (
	"errors"
	"fmt"
)

func main() {
	//示例代码
	var frenchHello string                 // 声明变量为字符串的一般方法
	var emptyString = ""                   // 声明了一个字符串变量，初始化为空字符串
	no, yes, maybe := "no", "yes", "maybe" // 短变量声明(简短声明)，同时声明多个变量
	japaneseHello := "Konichiwa"           // 同上
	frenchHello = "Bonjour"                // 常规赋值
	fmt.Println(frenchHello, emptyString)
	//"Bonjour"  ""
	fmt.Println(no, yes, maybe)
	//"no" "yes" "maybe"
	fmt.Println(japaneseHello)
	//"Konichiwa"
	var s = "hello"
	//s[0] = 'c' 无法对字符串中的值进行操作
	fmt.Println(s)
	//hello

	s1 := "hello"
	c := []byte(s1) // 将字符串 s 转换为 []byte 类型 (slice)
	c[0] = 'c'
	s2 := string(c) // 再转换回 string 类型
	fmt.Printf("%s\n", s2)
	//cello

	s3 := "hello,"
	m := " world"
	a := s3 + m
	fmt.Printf("%s\n", a)
	//hello,world
	fmt.Println(s3)
	//hello,

	s4 := "hello"
	s4 = "c" + s[1:] // 字符串虽不能更改，但可进行切片操作
	fmt.Printf("%s\n", s4)
	//cello

	m1 := `hello
    world`
	fmt.Println(m1)
	//hello
	//    world
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}

}
