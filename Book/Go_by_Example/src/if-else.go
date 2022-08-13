package main

import "fmt"

func main() {
	//判断奇偶
	if 7%2 == 0 {
		fmt.Println("7 is even") //偶数
	} else {
		fmt.Println("7 is odd") //奇数
	}
	//能否整除
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	//判断数是一位数还是多位数(0-∞),其余输出负数
	//注: C/C++ Java 都有的三元运算符，在Go中并没有
	//圆括号可以省略不写(推荐不写) {}是必需的
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
