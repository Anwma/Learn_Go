package main

import (
	"fmt"
	"runtime"
)

func main() {
	bool1 := true
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}
	//判断字符串是否为空
	str := ""
	if str == "" {
		//
	}
	if len(str) == 0 {
		//
	}

}

//判断操作系统的类型
var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isGreater(x, y int) bool {
	if x > y {
		return true
	}
	return false
}
