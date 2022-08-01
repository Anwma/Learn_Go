package main

import "fmt"

func trace2(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a2() {
	defer un(trace2("a2"))
	fmt.Println("in a2")
}

func b2() {
	defer un(trace2("b2"))
	fmt.Println("in b2")
	a2()
}

func main() {
	b2()
}
