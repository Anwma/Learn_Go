package main

import "fmt"

func main() {
	var i = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
