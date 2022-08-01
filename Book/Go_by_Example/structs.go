package main

import "fmt"

//创建结构体 person 包含name age两个字段
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//对结构体指针使用 . 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)
	//修改结构体对应的字段值
	sp.age = 51
	fmt.Println(sp.age)
}
