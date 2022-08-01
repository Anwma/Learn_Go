package main

import (
	"fmt"
	"os"
)

func main() {
	var s, str string
	for i := 0; i < len(os.Args); i++ {
		s += str + os.Args[i]
		str = " "
	}
	fmt.Println(s)
}
