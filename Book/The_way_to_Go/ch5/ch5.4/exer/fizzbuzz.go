package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch true {
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}

}
