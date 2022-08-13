package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	//支持多返回值
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
