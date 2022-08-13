package main

import "fmt"

func main() {
	array := []int{1, 2}
	length := len(array) + 1
	slice := make([]int, length)
	fmt.Println(slice)
	for i := 0; i < length-1; i++ {
		slice = append(slice,1)
		slice[array[i]]++
	}
}
