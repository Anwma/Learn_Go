package main

import "fmt"

func main() {
	fmt.Println(fibonacci(0))
}

func fibonacci(n int) int {
	//0 1 1 2 3 5
	switch n {
	case 0:
		return 0
	case 1, 2:
		return 1
	}
	q, p := 0, 1
	for i := 2; i < n; i++ {
		sum := q + p
		q = p
		p = sum
	}
	return p
}
