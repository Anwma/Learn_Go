package main

import (
	"fmt"
	"strconv"
)

type Stringer interface {
	String() string
}

type Human1 struct {
	name  string
	age   int
	phone string
}

// 通过这个方法 Human1 实现了 fmt.Stringer
func (h Human1) String() string {
	return "❰" + h.name + " - " + strconv.Itoa(h.age) + " years -  ✆ " + h.phone + "❱"
}

func main() {
	Bob := Human1{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human1 is : ", Bob)
}
