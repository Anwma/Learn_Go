package main

import "fmt"

func main() {
	f1("2", "34", "214")
	f2("2", "34", "214")
	f4(1, true, Options{var2: 2.2})
}

//写一个函数，该函数接收一个变长参数并对每个元素进行换行打印
//一个接受变长参数的函数可以将这个参数作为其他函数的参数进行传递
func f1(str ...string) {
	for _, i := range str {
		fmt.Println(i)
	}
	f2(str...)
	f3(str)
}
func f2(str ...string) {

}
func f3(str []string) {

}
func f4(a int, b bool, c Options) {

}

//如果变长参数的类型并不都相同，我们似乎要使用五个函数和参数来进行传递
//有以下两种方案可以解决这个问题

//1. 使用结构

type Options struct {
	var1 int
	var2 float64
	var3 bool
	var4 complex128
}

//2. 使用空接口

//func type-check(a int, b bool, values ...interface{}) {
//	for _, value := range values {
//		switch v := value.(type) {
//		case int:
//		case float64:
//		case string:
//		case bool:
//		default:
//		}
//	}
//}
