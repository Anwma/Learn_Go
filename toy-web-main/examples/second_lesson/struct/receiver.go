package main

import "fmt"

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	fmt.Printf("原始值：%v \n", u)
	u.ChangeName("Tom Changed!") //不变
	u.ChangeAge(100)             //修改
	fmt.Printf("修改值-1：%v \n", u)
	//{Tom 100}

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)

	fmt.Printf("修改值-2：%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// ChangeName 结构体接收器
func (u User) ChangeName(newName string) {
	u.Name = newName
}

// ChangeAge 指针接收器
func (u *User) ChangeAge(newAge int) {
	u.Age = newAge
}

type Handle func()

func (h Handle) Hello() {

}
