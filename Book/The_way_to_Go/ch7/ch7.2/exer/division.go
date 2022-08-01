package main

import "fmt"

func main() {
	buf := []byte{50}
	buf = []byte("hello-world")
	//分割buf 成两个切片
	s1, s2 := buf[0:5], buf[5:]
	fmt.Printf("s1:%c\ns2:%c\n", s1, s2)
}
