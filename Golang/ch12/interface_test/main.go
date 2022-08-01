package main

import (
	"fmt"
)

// Programmer 接口是一个协议 - 程序员 - 只要你能够1.写代码 2.解决bug 其实就是一组方法的集合
type Programmer interface {
	Coding() string //方法只是声明
	Debug() string
}
type Designer interface {
	Design() string
}
type Manager interface {
	Programmer
	Designer
	Manage() string
}
type G struct {
}

type Pythoner struct {
	UIDesigner
	lib   []string
	kj    []string
	years []int
}

//  Java的话 java里面一种类型只要继承一个接口才行 如果你继承了这个接口的话 那么这个接口里面的所有方法你必须要全部实现

type UIDesigner struct {
}

func (d UIDesigner) Design() string {
	fmt.Println("我会UI设计")
	return "我会UI设计"
}

func (p G) Coding() string {
	fmt.Println("go开发者")
	return "go开发者"
}
func (p G) Debug() string {
	fmt.Println("我会 go 的debug")
	return "我会go的debug"
}

func (p Pythoner) Coding() string {
	fmt.Println("python开发者")
	return "python开发者"
}
func (p Pythoner) Debug() string {
	fmt.Println("我会python的debug")
	return "我会python的debug"
}

func (p Pythoner) Manage() string {
	fmt.Println("不好意思,管理我也懂")
	return "不好意思,管理我也懂"
}

//func (p Pythoner) Design() string {
//	fmt.Println("我是一个python开发者,但是我会搞UI设计")
//	return "我是一个python开发者,但是我会搞UI设计"
//}

//对于 Pythoner 这个结构体来说 你实现任何方法都可以 但是你只要不全部实现这两个方法 那么你Pythoner结构体就不是一个Programmer类型
//1.Pythoner本身自己就是一个类型 那我何必在意我是不是Programmer
//2.封装 继承 多态 -多态的概念对于很多 Pythoner来说会有点陌生
//3.在讲解多态之前，我们来对interface做一个说明：在go语言中接口是一种类型 是一种抽象类型
func HandlePy(pythoner Pythoner) {

}
func HandleGo(g G) {

}

//func HandleJava(pythoner Pythoner) {
//
//}
type myError struct {
}

func (m myError) Error() string {
	return "错误"
}
func main() {
	//新的语言出来了 接口帮我们完成了go语言的多态
	//var pro Programmer = Pythoner{}
	//var g G = G{}
	//pro.Coding()
	var pros []Programmer
	pros = append(pros, Pythoner{})
	pros = append(pros, G{})

	//接口虽然是一种类型 但是和其他类型不太一样 接口是一种抽象类型 struct 是具象
	p := Pythoner{}
	fmt.Printf("%T\n", p)
	var pro Programmer = Pythoner{}
	var pro2 Programmer = G{}
	fmt.Printf("%T\n", pro)
	fmt.Printf("%T\n", pro2)
	//如果大家对于面向对象理解的话 java 里面的抽象类型有点像
	//1. go struct 组合 组合一起实现了所有的接的方法也是可以的
	//2.接口本身也支持组合

	var m Manager = Pythoner{}
	//struct 组合完成了接口 1.接口支持组合 - 继承2.结构体组合实现了所有的接口方法也没有问题
	m.Design()
	//python语言本身设计上是采用了完全的基于鸭子类型 - 协议 影响了python语法的for len()
	//go语言也推荐鸭子类型 error
	//var err error = error() //并不能实例化 抽象类型
	//var err error = myError{}
	//var err error = errors.New("错误")
	s := "文件不存在"
	var err error = fmt.Errorf("错误：%s", s)
	fmt.Println(err)
}

//开发中经常会遇到的问题
//开发一个电商网站 支付环境 wx zfb yhk 你的系统支持各种类型的支付 每一种支付类型都有统一的接口
//定一个协议 1. 支付 2.创建订单 3.查询支付状态 4.退款
//支付发起了

//type AliPay struct {
//}
//
//type WeChat struct {
//}
//type Bank struct {
//}
//
//var b Bank
//var a AliPay
//var w WeChat

type Tongyong struct {
}

//多态 声明类型的时候你声明的类型是一种兼容类型,但是实际赋值的时候是另一种类型
//接口是强制性
//你现在有一个缓存系统 - 这个地方你一开始使用的缓存是redis 但是后期你考虑到可能使用其他的缓存技术 - 本地 memcache

/*
	var x Tongyong
	x = Bank{}
	x = AliPay{}
	x = WeChat{}
	x.pay
	x.create
	x.query
*/
//这种多态特性 其实在python中不需要多态 python是动态语言
//go语言中并不支持继承
//如果后期接入一种新的支付 或者取消已有的支付
