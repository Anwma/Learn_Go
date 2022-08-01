package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "asSASA ddd dsjkdsjs dk"
	s2 := "asSASA ddd dsjkdsjsこん dk"
	//统计字节数
	fmt.Println(len(s1))
	fmt.Println(len(s2))
	//统计字符(rune)数
	fmt.Println(utf8.RuneCount([]byte(s1)))
	fmt.Println(utf8.RuneCount([]byte(s2)))
}
