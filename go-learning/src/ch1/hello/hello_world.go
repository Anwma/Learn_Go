package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
	//os.Exit(-1)
	//return 1
	//	func main must have no arguments and no return values
}
