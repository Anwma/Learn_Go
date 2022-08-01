package main

import (
	"fmt"
	"strings"
)

// 重复字符串
func main() {
	var origS = "Hi there! "
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}
