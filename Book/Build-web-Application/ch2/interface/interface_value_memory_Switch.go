package main

import (
	"fmt"
	"strconv"
)

type Element1 interface{}
type List1 []Element1

type Person1 struct {
	name string
	age  int
}

//打印
func (p Person1) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	List1 := make(List1, 3)
	List1[0] = 1       //an int
	List1[1] = "Hello" //a string
	List1[2] = Person1{"Dennis", 70}

	for index, Element1 := range List1 {
		switch value := Element1.(type) {
		case int:
			fmt.Printf("List1[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("List1[%d] is a string and its value is %s\n", index, value)
		case Person1:
			fmt.Printf("List1[%d] is a Person1 and its value is %s\n", index, value)
		default:
			fmt.Println("List1[%d] is of a different type", index)
		}
	}
}
