package main

import "fmt"

type person struct {
	name string
	age  int
}
type person1 struct {
	name string
	age  int
}

func main() {
	var P person1 // P现在就是person类型的变量了

	P.name = "Astaxie"                            // 赋值"Astaxie"给P的name属性.
	P.age = 25                                    // 赋值"25"给变量P的age属性
	fmt.Printf("The person's name is %s", P.name) // 访问P的name属性.
}
