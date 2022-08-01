package main

import (
	"log"
)

func main() {
	//where := func() {
	//	/*
	//		    调用方报告关于调用goroutine堆栈上函数调用的文件和行号信息。
	//			参数skip是上升的堆栈帧数，0标识调用者的调用者。(由于历史原因，
	//			Caller和Caller之间skip的含义不同。)返回值报告相应调用的文件中
	//			的程序计数器、文件名和行号。如果无法恢复信息，则布尔值ok为假。
	//	*/
	//	_, file, line, _ := runtime.Caller(1)
	//	log.Printf("%s:%d", file, line)
	//}
	//where()
	//// some code
	//where()
	//// some more code
	//where()
	func1()

}

var where = log.Print

func func1() {
	where()

	where()

	where()
}
