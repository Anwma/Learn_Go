package main

import "fmt"

//创建map 类型的切片
func main() {
	// Version A: 创建一个k-int v-int size=5 的map 命名为items
	items := make([]map[int]int, 5)
	for i := range items {
		//items中每一个键对应一个 map key-int v-int size=1
		items[i] = make(map[int]int, 1)
		//items[i] 代表 items中每一个元素
		//items[i][j] 代表 items中每一个元素 中所对应的值
		items[0][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
