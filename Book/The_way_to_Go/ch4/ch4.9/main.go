package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++
	fmt.Println(a)
}
