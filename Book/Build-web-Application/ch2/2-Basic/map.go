package main

import "fmt"

func main() {
	// 声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	//var numbers map[string]int

	// 另一种map的声明方式
	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值0
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("numbers: ", numbers)
	fmt.Println("第三个数字是: ", numbers["three"]) // 读取数据
	// 打印出来如:第三个数字是: 3

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Println(rating)
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // 删除key为C的元素

	m := make(map[string]string)
	fmt.Println("m: ", m)
	m["Hello"] = "Bonjour"
	fmt.Println("m: ", m)
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了
	fmt.Println("m1: ", m1)

}
