package main

import "fmt"

func main() {
	a, b, c := hechaji1(2, 3)
	fmt.Println(a, b, c)
	d, e, f := hechaji2(2, 3)
	fmt.Println(d, e, f)

}
func hechaji1(a int, b int) (int, int, int) {
	return a + b, a * b, a - b
}
func hechaji2(a1 int, b1 int) (d int, e int, f int) {
	d = a1 + b1
	e = a1 * b1
	f = a1 - b1
	return d, e, f
}
