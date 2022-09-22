package main

import (
	"day1/util"
	"fmt"
	"runtime"
	"strconv"
)

func atc() {
	var a float32 = 8
	var b float32 = 3
	var c float32 = a + b
	fmt.Printf("c=%f\n", c)
	var d float32 = a - b
	var e float32 = a * b
	var f float32 = a / b
	var g float32 = 10 % 3
	fmt.Println(d, e, f, g)
}

func rel() {
	a, b, c := 8, 3, 5
	fmt.Println(a > b && b > c)
	fmt.Println(a > b || b > c)
	fmt.Println(!(b > c))
}

func bitOp() {
	fmt.Printf("os arch %s, int size %d\n", runtime.GOARCH, strconv.IntSize)
	fmt.Printf("260 %s\n", util.BinaryFormat(260))
	fmt.Printf("260 %s\n", util.BinaryFormat(260>>11))
	fmt.Printf("-260 %s\n", util.BinaryFormat(-260))
	fmt.Printf("-260 %s\n", util.BinaryFormat(-260<<23)) //左移 右侧补0
	fmt.Printf("-260 %s\n", util.BinaryFormat(-260>>23)) //右移 右侧补符号位
	fmt.Printf("%d\n", ^5)

}
func varType() {
	a := 14641313214.2
	fmt.Printf("%e", a)
}

const (
	aaa = iota
	bbb = 30
	ccc = iota
	ddd
)
const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
)
const (
	aa, bb = iota + 1, iota + 2 //1 2 iota=0 aa=1 bb=2
	cc, dd                      //2 3 iota=1 cc集成aa=iota+1->1+1=2
	ee, ff                      //3 4
)

func main() {
	fmt.Println(aaa, bbb, ccc, ddd)
	fmt.Println(kb, mb, gb)
	fmt.Println(aa, bb, cc, dd, ee, ff)
	atc()
	rel()
	bitOp()
	varType()
	fmt.Printf("%t\n", 04 == 4.00) //用到了整形字面量和浮点型字面量
	fmt.Printf("%v\n", .4i)
	fmt.Printf("%t\n", '\u4f17' == '众')
	a := 2
	ptr := &a
	fmt.Printf("%d\n", ptr) //地址的十进制表示
	fmt.Printf("%p\n", ptr) //地址的16进制表示
	fmt.Printf("%p\n", &a)  //

}
