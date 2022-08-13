package main

import (
	"fmt"
	"os"
)

//获取命令行参数
func main() {
	s,sep := "",""
	for _, arg := range os.Args[1:] {
		//os.Args[:]的类型是[]string,采取遍历切片的方式
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	//fmt.Println(os.Args[1:]) //[-test hello]
}
