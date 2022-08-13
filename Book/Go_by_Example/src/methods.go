package main

import "fmt"

type rect1 struct {
	width, height int
}

//method1
func (r *rect1) area() int {
	return r.width * r.height
}

//method2
func (r rect1) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	//rect是矩形
	//创建结构体r(rect)并赋默认值
	r := rect{width: 10, height: 5}
	//使用m1计算面积
	fmt.Println("area: ", r.area())
	//使用m1计算周长
	fmt.Println("perim:", r.perim())
	//通过使用指针来避免拷贝(浪费空间)
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
