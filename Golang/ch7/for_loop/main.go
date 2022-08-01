package main

import "fmt"

func main() {
	//for init; condition; post{} 计算1-10的和
	//sum := 0
	//for i := 1; i <= 10; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//1. i:=1
	//2. i是否 <=10
	//3. sum = 1
	//4. i++ ->i =2
	//5. ...
	//6. i=11 跳出循环

	//循环字符串
	name := "Wozen:我的世界"
	//for index, value := range name {
	//fmt.Println(index, value)
	//	fmt.Printf("%d %c\n", index, value)
	//}

	//1. name 是一个字符
	//2. 字符串是字符串的数组
	//fmt.Printf("%c\n", name[0])
	//fmt.Printf("%c\n", name[6])
	name_arr := []rune(name)
	for i := 0; i < len(name_arr); i++ {
		fmt.Printf("%c\n", name_arr[i])
	}
	//3.在做字符串遍历的时候，要尽量使用range
}
