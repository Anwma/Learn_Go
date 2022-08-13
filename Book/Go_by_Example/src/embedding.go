package main

import "fmt"

//创建base结构体
type base struct {
	age int
	num int
}

//(参数列表):b base (返回值类型)describe() string
func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

//创建container结构体，包含base?
type container struct {
	base
	str string
}

func main() {
	//co 是container的一个对象
	co := container{
		//必须对base进行显示初始化
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	//co(container)->base str->base.describe
	//两种访问方法均可
	//fmt.Println("describe:", co.base.describe()) --1
	fmt.Println("describe:", co.describe()) //--2
	//创建一个describer接口 访问describe方法
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}

// Circle ----------------------------
//比如说，在2D绘图程序当中，我们可能会定义以下两个结构类型
//圆圈
/*
type Circle struct {
	X, Y, Radius int
}
//Wheel 车轮
type Wheel struct {
	X, Y, Radius, Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}
*/
//--------------------------------------
//但是当支持的形状变多了之后 我们需要重构相同的部分
//--------------------------------------

/*
type Point struct {
	X,Y int
}
type Circle struct {
	Center Point
	Radius int
}
type Wheel struct {
	Circle Circle
	Spokes int
}

func main()  {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}
*/
//程序结构简单了，但是访问wheel成员变得麻烦了
//----------------------------
//结构体嵌套和匿名成员的引入
//匿名成员：Go允许我们定义不带名称的结构体成员，只需要指定类型即可；
//这种结构体成员称作匿名成员。这个结构体成员的类型必须是一个命名类型
//或者指向命名类型的指针
/*
type Point struct {
	X,Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func main()  {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}
*/
