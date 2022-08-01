package main

import (
	"fmt"
	"os"
)

func main() {
	s, str := "", ""
	for _, arg := range os.Args[1:] {
		s += str + arg
		str = " "
	}
	fmt.Println(s)
}
