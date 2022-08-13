package main

import (
	"fmt"
	"math"
)

//接口对area perim 两种方法进行了再次抽象
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

//geometry:几何体
func measure(g geometry) {
	//输出自身: {3 4}、 {5}
	fmt.Println(g)
	//输出面积/周长
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//创建r(rest)结构体 宽3高4
	r := rect{width: 3, height: 4}
	//创建⚪(circle) 半径5
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
