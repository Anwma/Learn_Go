package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		var t int
		for j := i; j != 0; j /= 10 {
			t = t*10 + j%10
		}
		if t == i {
			fmt.Printf("%d\n", i)
		}
	}
}
