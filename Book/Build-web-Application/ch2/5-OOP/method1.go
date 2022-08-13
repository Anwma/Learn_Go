package main

import (
	"fmt"
	"math"
)

type Rectangle1 struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle1) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//type  类型名 类型字面文字
//type typeName typeLiteral

func main() {
	r1 := Rectangle1{12, 2}
	r2 := Rectangle1{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of r2 is: ", r2.area())
	fmt.Println("Area of c1 is: ", c1.area())
	fmt.Println("Area of c2 is: ", c2.area())

	type ages int

	type money float32

	type months map[string]int

	m := months{
		"January":  31,
		"February": 28,
		//...
		"December": 31,
	}
	fmt.Println(m)
}
