package main

func IsValidTriangle(a int, b int, c int) bool {
	// write your code here
	return a+b > c && a+c > b && b+c > a
}
