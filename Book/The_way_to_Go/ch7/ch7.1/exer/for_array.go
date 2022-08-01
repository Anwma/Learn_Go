package main

import "fmt"

func main() {
	n := 16
	arr := [16]int{}
	for i := 0; i < n; i++ {
		arr[i] = i
		fmt.Println(arr[i])
	}
}
