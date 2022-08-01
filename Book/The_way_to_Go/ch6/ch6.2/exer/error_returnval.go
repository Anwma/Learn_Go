package main

import "fmt"

func main() {
	fmt.Println(MySqrt(2))
	fmt.Println(MySqrt2(2))
}
func MySqrt(a float64) float64 {
	return a * a
}
func MySqrt2(a float64) (sqrt float64) {
	sqrt = a * a
	return sqrt
}
