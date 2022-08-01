package main

import "fmt"

func main() {
	//0 1 2 3 4 5
	//1 1 2 3 5 8
	fmt.Println(fb(10))
	for i := 10; i > 0; i-- {
		r := dg(i)
		fmt.Println(r)
	}
	// 使用int 类型最多只能计算到 12 的阶乘，
	// 因为一般情况下 int 类型的大小为 32 位，
	// 继续计算会导致溢出错误。
	for i := 12; i > 0; i-- {
		r := fact(i)
		fmt.Println(r)
	}
}

//练习6.4
func fb(n int) (p int, res int) {
	if n <= 1 {
		p, res = 1, 1
	} else {
		_, v1 := fb(n - 1)
		_, v2 := fb(n - 2)
		p, res = n, v1+v2
	}
	return
}

//练习6.5
func dg(n int) (r int) {
	return n
}

//练习6.6
func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}
