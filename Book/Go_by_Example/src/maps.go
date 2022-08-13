package main

import "fmt"

func main() {

	m := make(map[string]int)
	//建立map key-value类型 key:string value:int
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m)) //2
	//删除key 为 k2的键
	delete(m, "k2")
	fmt.Println("map:", m)
	//面对不需要的值用 _ 空白标识符
	_, prs := m["k2"]
	fmt.Println("prs:", prs) //false
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
