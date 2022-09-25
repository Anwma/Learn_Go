package main

//
//import (
//	"strings"
//)
//
//func Solution(n int, s string) int {
//	// Write your code here.
//	//corner case
//	if s == "" {
//		return 2 * n
//	}
//	//创建一个二维切片
//	slice := make([][]int, n)
//	for i := 0; i < n; i++ {
//		slice[i] = make([]int, 10)
//	}
//	//对s中给定的字符串进行切分 转换成切片
//	str := strings.Split(s, " ") //str = ["1A" "2B" "1C"]
//	for _, v := range str {
//		slice[v[0]][v[1]-65] = 1
//	}
//	return 1
//}
//
//func main() {
//	Solution(2,"1A 2B 1C")
//}
