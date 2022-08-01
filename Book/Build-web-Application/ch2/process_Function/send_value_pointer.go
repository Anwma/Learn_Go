package main

import "fmt"

//简单的一个函数，实现了参数+1的操作
func add1(a int) int {
	a = a + 1 // 我们改变了a的值
	return a  //返回一个新值
}

//简单的一个函数，实现了参数+1的操作
func add2(a *int) int { // 请注意，
	*a = *a + 1 // 修改了a的值
	return *a   // 返回新值
}
func main() {
	x := 3
	_x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := add1(x)   //调用add1(x)
	x2 := add2(&_x) //调用add2(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"
	fmt.Println("----------------------------")
	fmt.Println("_x+1 = ", x2) // 应该输出"x = 4"
	fmt.Println("_x = ", _x)   // 应该输出"_x = 4"

}
