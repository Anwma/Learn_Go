package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 5, 6, 7}
	var b = []int{8, 9, 10, 11, 12, 13}
	//1.将切片 b 的元素追加到切片 a 之后:
	a = append(a, b...)
	//2.复制切片 a 的元素到新的切片 b 上:
	b = make([]int, len(a))
	copy(b, a)
	//3.删除位于索引 i 的元素:
	a = append(a[:2], a[3:]...)
	//4.切除切片 a 中从索引 i 至 j 位置的元素：
	a = append(a[:1], a[3:]...)
	//5.为切片 a 扩展 j 个元素长度
	a = append(a, make([]int, 5)...)
	//6.在索引 i 的位置插入元素 x:
	a = append(a[:3], append([]int{2, 3}, a[3:]...)...)
	//7.在索引 i 的位置插入长度为 j 的新切片：
	a = append(a[:2], append(make([]int, 3), a[2:]...)...)
	//8.在索引 i 的位置插入切片 b 的所有元素：
	a = append(a[:3], append(b, a[3:]...)...)
	//9.取出位于切片 a 最末尾的元素 x：
	x := make([]int, 1)
	x, a = a[len(a)-1:], a[:len(a)-1]
	fmt.Println(x)
	//10.将元素 x 追加到切片 a
	x1 := 2
	a = append(a, x1)
}
