package main

import (
	"fmt"
	"os"
	"regexp"
)

//未能正确读到数据 暂时放着
func main() {
	a := FindFileDigits("wx.txt")
	fmt.Println(a)
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindFileDigits(filename string) []byte {
	fileBytes, _ := os.ReadFile(filename)
	b := digitRegexp.FindAll(fileBytes, len(fileBytes))
	c := make([]byte, 1)
	for _, bytes := range b {
		c = append(c, bytes...)
	}
	return c
}
