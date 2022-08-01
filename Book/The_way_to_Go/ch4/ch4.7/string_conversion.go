package main

import (
	"fmt"
	"strconv"
)

// 字符串与其它类型的转换
func main() {
	var orig = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	//若orig 不能转换为整数，则an = 0, _ 忽略err 信息
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
