package main

import (
	"fmt"
	"os"
)

func main() {
	//输出命令行参数
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	/*
		PS D:\JetbrainsCode\Learn_Go\Book\Black_Book\ch1\echo> go run .\echo1.go -hello
		C:\Users\Anwma\AppData\Local\Temp\go-build40468402\b001\exe\echo1.exe -hello
	*/
}
