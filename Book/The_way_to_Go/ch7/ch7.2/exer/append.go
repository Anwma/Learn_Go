package main

import "fmt"

func main() {
	sl := []byte("You")
	res := Append(sl, []byte("23"))
	fmt.Printf("%c", res)
}
func Append(src []byte, data []byte) (res []byte) {
	res = append(src, data...)
	return res
}
