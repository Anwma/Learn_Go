package main

import "fmt"

func main() {
	s := "HelloWorld!"
	s1, s2 := split(s, 5)
	fmt.Println(s1 + "----" + s2)
}

//编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，
//然后返回两个分割后的字符串。
func split(s string, index int) (string, string) {
	str := []byte(s)
	slice1, slice2 := str[0:index], str[index:]
	return string(slice1), string(slice2)
}
