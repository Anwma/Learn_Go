package main

import "fmt"

func main() {
	var s = []int{1, 2}
	dst := increbyfactor(s, 3)
	fmt.Printf("源slice-len:%d\n新的slice-len:%d", len(s), len(dst))
}
func increbyfactor(src []int, n int) (dst []int) {
	newSlice := make([]int, len(src)*n)
	copy(newSlice, src)
	return newSlice
}
