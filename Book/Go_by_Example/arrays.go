package main

import "fmt"

func main() {
	//创建int型数组a,大小为5,数组内默认值为0
	var a [5]int
	fmt.Println("emp:", a)
	//对int数组a的索引为4的元素赋值为100
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	//判定长度
	fmt.Println("len:", len(a))
	//采用短变量声明一个int数组b,并指定相应位置的值

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	//与b的创建一样，初始为空数组，大小5
	c := [5]int{}
	fmt.Println("c:", c)
	//创建二维数组 默认为空
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
