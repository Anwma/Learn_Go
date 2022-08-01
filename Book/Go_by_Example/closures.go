package main

import "fmt"

func intSeq() func() int {
	//函数式编程，可以参考《计算机程序构造与解释》一书(Lisp)
	i := 0 //----1
	return func() int {
		i++      //----2
		return i //----3
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
