package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func scope() {
	var a int32
	var pointer unsafe.Pointer = unsafe.Pointer(&a)
	var p uintptr = uintptr(pointer)
	var ptr *int32 = &a
	fmt.Printf("pointer %p\n", pointer)
	fmt.Printf("%v\n", p)
	fmt.Printf("%p\n", ptr)
}

type User struct {
	Name string
	Age  int
}

func (self User) Hello() {
	fmt.Printf("my name is %s , age is %d", self.Name, self.Age)
}

type ms map[string]int32

func (self ms) Say() {
	fmt.Printf("%d\n", self["hello"])
}

type s uint8

func main() {
	//var a s
	//a = 1
	//fmt.Print(a)
	//scope()
	//Jeff := User{
	//	Name: "Jeff",
	//	Age:  12,
	//}
	//Jeff.Hello()
	ddd := "123425"
	text := strings.Split(ddd, "")
	fmt.Println(text)

	res := strings.Contains(ddd,"2")
	fmt.Println(res)


	tmp := strings.HasPrefix(ddd,"12")
	fmt.Println(tmp)

	//后缀
	tmp2 := strings.HasSuffix(ddd,"5")
	fmt.Println(tmp2)

	tmp3 := strings.Index(ddd,"2")
	//最后一次出现的位置
	//tmp4 := strings.LastIndex(ddd,"2")
	//tmp5 := strings.Count(ddd,"2")

	fmt.Println(tmp3)
	//fmt.Println(tmp4)
	//fmt.Println(tmp5)
}
