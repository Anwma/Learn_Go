package main

import "fmt"

func main() {
	ch := "G"
	for i := 25; i > 0; i-- {
		for j := i; j < 26; j++ {
			fmt.Print(ch)
		}
		fmt.Println()
	}

}
