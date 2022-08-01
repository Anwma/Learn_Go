package main

import "fmt"

// 对一个空指针的反向引用是不合法的，并且会使程序崩溃
func main() {
	var p *int = nil
	//*p = 0 不能够对空指针进行赋值
	fmt.Println(p)
}

// in Windows: stops only with: <exit code="-1073741819" msg="process crashed"/>
// runtime error: invalid memory address or nil pointer dereference
