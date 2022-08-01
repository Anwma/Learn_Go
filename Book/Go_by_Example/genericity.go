package main

import "fmt"

// MapKeys 泛型：类型参数
// -----------------------------------------------------------------
// type comparable interface{ comparable }
// comparable是一个由所有可比较类型实现的接口
// (布尔值、数字、字符串、指针、通道、可比类型数组、结构的字段都是可比较的类型)。
// 可比接口只能作为类型参数约束使用，不是变量的类型。
// -----------------------------------------------------------------
// type any = interface{}
// any是接口{}的别名，在所有方面都等价于接口{}。
func MapKeys[K comparable, V any](m map[K]V) []K {
	//属实没看懂...
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	//创建一个map key(int)-value(string) 并赋初值
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	//打印keys值 1 2 4
	fmt.Println("keys m:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
