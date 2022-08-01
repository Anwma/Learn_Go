package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	fmt.Printf("%b", &arr1)
	fmt.Printf("%b", &arr2)
}
