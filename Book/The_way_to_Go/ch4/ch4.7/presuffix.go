package main

import (
	"fmt"
	"strings"
)

// 判断某个字符串的前缀
// strings.HasPrefix(str,"")
func main() {
	var str = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
