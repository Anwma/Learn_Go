package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 3, 4, 1, 0, 6, 9, 7, 8}
	fmt.Println("排序前: ", a)
	sort.Ints(a)
	fmt.Println("排序后: ", a)
}
