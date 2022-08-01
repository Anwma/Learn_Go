package main

import "fmt"

func main() {
	var s = "yanandaxue"
	str := []byte(s)
	//翻转字符串 pos: len(str/2)
	str1 := append(str[len(str)/2:], str[:len(str)/2]...)
	fmt.Printf("%c", str1)
}
