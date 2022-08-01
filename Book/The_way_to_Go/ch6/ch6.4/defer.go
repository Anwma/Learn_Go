package main

import "fmt"

func main() {
	//function1()
	//a()
	//f()
}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

//func a() {
//	i := 0
//	defer fmt.Println(i)
//	i++
//	return
//}
//入栈 0 1 2 3 4
//出栈 4 3 2 1 0 释放defer
//func f() {
//	for i := 0; i < 5; i++ {
//		defer fmt.Printf("%d ", i)
//	}
//
//}

//defer 允许我们进行一些函数执行完成后的收尾工作
// open a file
//defer file.Close()

//mu.Lock()
//defer mu.Unlock()

//printHeader()
//defer printFooter()

// open a database connection
//defer disconnectFromDB()
