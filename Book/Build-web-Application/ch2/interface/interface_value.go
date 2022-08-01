package main

import "fmt"

type Humans struct {
	name  string
	age   int
	phone string
}

type Students struct {
	Humans //匿名字段
	school string
	loan   float32
}

type Employees struct {
	Humans  //匿名字段
	company string
	money   float32
}

// SayHi Humans实现SayHi方法
func (h Humans) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing Humans实现Sing方法
func (h Humans) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// SayHi Employees重载Humans的SayHi方法
func (e Employees) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Mens Interface Mens被Humans,Students和Employees实现
// 因为这三个类型都实现了这两个方法
type Mens interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Students{Humans{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Students{Humans{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employees{Humans{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employees{Humans{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//定义Mens类型的变量i
	var i Mens

	//i能存储Students
	i = mike
	fmt.Println("This is Mike, a Students:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employees
	i = tom
	fmt.Println("This is tom, an Employees:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Mens
	fmt.Println("Let's use a slice of Mens and see what happens")
	x := make([]Mens, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
