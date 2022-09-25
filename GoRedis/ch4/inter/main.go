package main

//
//import (
//	"fmt"
//)
//
//type Car interface {
//	Driver()
//}
//type TrafficTool interface {
//	Driver()
//	Run()
//}
//type Truck struct {
//	Model string
//}
//
//func (t Truck) Driver() {
//	fmt.Println(t.Model)
//}
//
//func main() {
//	//1.一个接口的值在底层是如何表示的
//	//c的底层表示是iface
//	//data 指针指向的才是我们新建的结构体
//	//var c Car = Truck{}
//	var c Car = &Truck{}
//	//truck := c.(Truck)
//	//tool := c.(TrafficTool)
//	fmt.Printf("%p\n", c)
//	//fmt.Printf("%+p\n", truck)
//	//fmt.Printf("%p\n", tool)
//
//	//switch c.(type) {
//	//case Truck:
//	//	fmt.Println("Truck")
//	//case Car:
//	//	fmt.Println("Car")
//	//case TrafficTool:
//	//	fmt.Println("TrafficTool")
//	//}
//}
