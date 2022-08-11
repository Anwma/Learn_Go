// 7.12
package main

import "fmt"

func main() {
	var s string = "hello,world"
	//str := []byte(s)
	//fmt.Printf("%T", str)
	s1, s2 := stringSplit(s, 2)
	fmt.Println(s1)
	fmt.Println(s2)
}

func stringSplit(s string, i int) (string, string) {
	str := []byte(s)
	return string(str[:i]), string(str[i:])
}
