package main

import (
	"fmt"
	"os"
)

const (
	i      = 100
	pi     = 3.1415
	prefix = "Go_"
)

var (
	i1      int
	pi1     float32
	prefix1 string
)

func main() {
	fmt.Println()
	fmt.Println(i, pi, prefix)
	fmt.Println(i1, pi1, prefix1)
	defer os.Exit(0)
	/*

		int     0
		int8    0
		int32   0
		int64   0
		uint    0x0
		rune    0 //rune的实际类型是 int32
		byte    0x0 // byte的实际类型是 uint8
		float32 0 //长度为 4 byte
		float64 0 //长度为 8 byte
		bool    false
		string  ""

	*/
}
