package main

import "fmt"

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	fmt.Println(s)
	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	fmt.Println(s)

	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)
	fmt.Println(s)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)
	fmt.Println(s)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)
	fmt.Println(s)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)
	fmt.Println(s)

}

func Add(s []int, index int, value int) []int {
	//TODO
	//在索引范围内,添加到中部(包括index = 0)
	if index < len(s) && index >= 0 {
		s = append(s, 0)
		copy(s[index+1:], s[index:])
		s[index] = value
	} else if index >= len(s) {
		//超出索引范围
		s = append(s, value)
	}

	return s
}

func Delete(s []int, index int) []int {
	// TODO
	if index < len(s) && index >= 0 {
		s = append(s[:index], s[index+1:]...)
	}
	return s
}
