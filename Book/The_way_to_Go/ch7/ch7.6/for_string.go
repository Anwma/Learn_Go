package main

import "fmt"

func main() {
	//s := "\u00ff\u754c"
	//for i, c := range s {
	//	fmt.Printf("%d:%c ", i, c)
	//}
	var b []byte
	var s string
	s = "123"
	b = append(b, s...)
	fmt.Printf("%c", b)
}
