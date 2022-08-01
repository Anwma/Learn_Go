package main

import "fmt"

func main() {
	/*

		func myfunc(arg ...int) {}

	*/
	arg := make([]int, 0)
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}
