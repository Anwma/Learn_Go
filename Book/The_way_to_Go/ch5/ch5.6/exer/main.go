package main

import "fmt"

func main() {
	//====================
	//Value of i is:0 1 2
	//A statement just after for loop.
	i := 0
	for { //since there are no checks, this is an infinite loop
		if i >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", i)
		i++
	}
	fmt.Println("A statement just after for loop.")
	//====================
	//1 3 5
	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}
}
