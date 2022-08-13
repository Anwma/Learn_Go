package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 3
	t := reflect.TypeOf(i)  //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i) //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值

	tag := t.Elem().Field(0).Tag       //获取定义在struct里面的标签
	name := v.Elem().Field(0).String() //获取存储在第一个字段里面的值
	fmt.Println(tag, name)

	var x float64 = 3.4
	v1 := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	fmt.Println(v1)

	var x1 float64 = 3.4
	v2 := reflect.ValueOf(x)
	v2.SetFloat(7.1)
	fmt.Println(x1)

	var x2 float64 = 3.4
	p := reflect.ValueOf(&x)
	v3 := p.Elem()
	v3.SetFloat(7.1)
	fmt.Println(x2)
}
