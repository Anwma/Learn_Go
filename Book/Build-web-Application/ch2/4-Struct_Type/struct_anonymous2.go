package main

import "fmt"

type Human2 struct {
	name  string
	age   int
	phone string // Human2类型拥有的字段
}

type Employee2 struct {
	Human2     // 匿名字段Human2
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	Bob := Employee2{Human2{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human2的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human2.phone)
}
