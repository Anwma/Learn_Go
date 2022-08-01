package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	// Note the explicit address-of operator
	x := Sum(&array)
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a { // dereference-ing *a to get back to the array is not necessary!
		sum += v
	}
	return
}