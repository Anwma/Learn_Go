package main

import "fmt"

func main() {
	n := 0
	for i := 0; i < 15; i++ {
		fmt.Printf("第 %d 个值为: %d \n", n, n)
		n += 1
	}

}
