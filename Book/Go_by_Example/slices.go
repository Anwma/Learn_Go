
package main

import "fmt"

func main() {
	//创建切片，默认值为空字符串 ""
	s := make([]string, 3)
	fmt.Println("emp:", s)
	//对s切片赋值a b c
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	//获取长度
	fmt.Println("len:", len(s))
	//向s中追加 "d"
	s = append(s, "d")
	//slice 扩容可查看src/runtime/slice.go 下的 growslice 函数
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	//创建与c同样大小的c slice
	c := make([]string, len(s))
	//切片复制
	copy(c, s)
	fmt.Println("cpy:", c)
	//左取右不取
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	//二维切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	//OutPut: [[0] [1 2] [2 3 4]]
}
