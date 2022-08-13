package main

import "fmt"

// SumAndProduct 返回 A+B 和 A*B
func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}
func SumAndProduct2(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}
func main() {
	x := 3
	y := 4

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
