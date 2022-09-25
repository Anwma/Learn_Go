package main

import (
	"fmt"
)

type Car interface {
	Driver()
}

type Truck struct {
	Model string
}

func (t Truck) Driver() {
	fmt.Println(t.Model)
}

func k(interface{}) {

}
func main() {
	//var c Car = &Truck{}
	//fmt.Println(c)
	k("123")
}
