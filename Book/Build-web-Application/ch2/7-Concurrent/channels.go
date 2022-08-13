package main

import "fmt"

// sum 对a中的切片元素求和,并将和返回到 chan 中
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	//开两个协程计算前半部分和后半部分的和
	go sum(a[:len(a)/2], c) //前半部分
	go sum(a[len(a)/2:], c) //后半部分
	/*
		x := <-c
		fmt.Println(x) //-5
		y := <-c
		fmt.Println(y) //17
	*/
	//采取就近原则
	x, y := <-c, <-c        // receive from c
	fmt.Println(x, y, x+y)
}
