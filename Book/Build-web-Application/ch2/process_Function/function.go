package main

import "fmt"

// 返回a、b中最大值.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	/*

		func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
		    //这里是处理逻辑代码
		    //返回多个值
		    return value1, value2
		}

	*/
	x := 3
	y := 4
	z := 5

	maxXy := max(x, y) //调用函数max(x, y)
	maxXz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, maxXy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, maxXz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它
}
