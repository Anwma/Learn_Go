package main

import "fmt"

func main() {
	//goto能不用就不用 goto过多 label过多 整个程序到后期维护就会麻烦
	//最容易理解的代码就是逐行的执行，哪怕你多一个函数的调用对于我们都是理解上的负担
	/* 定义局部遍历 */
	var a = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为：%d\n", a)
		a++
	}
}
