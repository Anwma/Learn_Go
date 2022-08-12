package main

import "fmt"

func main() {
	for i := 1000; i <= 1500; i++ {
		fmt.Printf("%d;", i)
		if i % 20 == 0{
			fmt.Println()
			fmt.Print("md ")
		}
	}
}
