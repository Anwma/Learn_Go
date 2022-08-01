package main

import "fmt"

func main() {
	fv := func(u string) { fmt.Println(u) }
	fv("Hello World")
	fmt.Printf("%T", fv)
}
