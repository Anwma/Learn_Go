package main

import "fmt"

func main() {
	//fmt.Print("md ")
	for i := 0; i <= 500; i++ {
		if i % 20 == 0 {
			fmt.Println()
			fmt.Print("md ")
		}
		fmt.Printf("%d;", i)
	}
	fmt.Println("\npause")
}
