package main

import (
	"fmt"
	"os"
)

func main() {
	//s,str := "",""
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, os.Args[i])
	}
}
