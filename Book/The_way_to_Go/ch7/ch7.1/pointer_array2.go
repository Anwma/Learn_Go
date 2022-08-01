package main

import "fmt"

func fp1(a *[3]int) { fmt.Println(a) }

func main() {
	for i := 0; i < 3; i++ {
		fp1(&[3]int{i, i * i, i * i * i})
	}
}
