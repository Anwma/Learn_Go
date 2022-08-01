package main

import "fmt"

//值传递
func zeroval(ival int) {
	ival = 0
}

//“引用”传递 ->本质还是值传递，只不过值变成了指针所指的逻辑地址
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	//打印i所在的逻辑地址(每个人机器打出来的值都会有所不同,由OS分配)
	//fmt.Println(&i)  0xc000016098
	fmt.Println("pointer:", &i)
}
