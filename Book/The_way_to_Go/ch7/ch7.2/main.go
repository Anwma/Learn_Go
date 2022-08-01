package main

import (
	"fmt"
)

func main() {
	//给定 s := make([]byte, 5)，len (s) 和 cap (s) 分别是多少？
	//s = s[2:4]，len (s) 和 cap (s) 又分别是多少？
	s := make([]byte, 5)
	fmt.Println(len(s), cap(s)) // 5 5
	s = s[2:4]
	fmt.Println(len(s), cap(s)) // 2 3

	//假设 s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := s1[2:]，s2 的值是多少？
	//如果我们执行 s2[1] = 't'，s1 和 s2 现在的值又分别是多少？
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:]
	fmt.Printf("%c", s2) //em
	//s2 是对s1 抽象 引用 slice 本身是引用类型。当对上层修改时，下层也会发生变化
	s2[1] = 't'
	fmt.Printf("%c", s1) // poet
	fmt.Printf("%c", s2) //et
	//读写长度未知的bytes最好使用buffer
	//var buffer bytes.Buffer
	//var r = new(bytes.Buffer)

	//通过buffer串联字符串
	//var buffer bytes.Buffer
	//for {
	//	if str, ok := getNextString(); ok { //method getNextString() not shown here
	//		buffer.WriteString(str)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Print(buffer.String(), "\n")
}
