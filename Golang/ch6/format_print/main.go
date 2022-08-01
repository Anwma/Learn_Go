package main

import "fmt"

func main() {
	//printf println
	//name := "Steven"
	//age := 18
	//fmt.Println("name:" + name + ",age:" + string(age))
	//fmt.Printf("name: %v, age: %v\n", name, age)
	//fmt.Printf("name: %#v, age: %#v\n", name, age)
	//fmt.Printf("name: %T, age: %T\n", name, age)
	//desc := fmt.Sprintf("name: %T, age: %T\n", name, age)
	//fmt.Println(desc)

	//data := 65
	//fmt.Printf("%c\n", data)
	//fmt.Printf("%q\n", data)
	//fmt.Printf("%U\n", data)
	//fmt.Printf("%#U\n", data)
	//
	//f := 3.1415926
	//fmt.Printf("%e\n", f)
	//fmt.Printf("%f\n", f)

	//输入
	var name string
	var age int
	//fmt.Println("请输入你的姓名和年龄: ")
	//fmt.Scanln(&name, &age)
	//fmt.Println(name, age)

	//通过scanf输入
	fmt.Println("请输入你的姓名和年龄: ")
	fmt.Scanf("%s %d", &name, &age) //按照原样的格式进行输入
	fmt.Println(name, age)
}
