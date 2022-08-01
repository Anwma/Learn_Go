package main

import (
	"fmt"
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
	c := 3 + 4i
	fmt.Println(real(c))
	fmt.Println(imag(c))
}
