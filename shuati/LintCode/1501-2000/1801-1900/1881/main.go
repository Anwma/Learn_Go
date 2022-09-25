package main

import (
	"strconv"
	"strings"
)

func Solution(n int, s string) int {
	// Write your code here.
	//corner case
	if s == "" {
		return 2 * n
	}
	//创建一个二维切片
	p := make([][]int, n)
	for i := 0; i < n; i++ {
		p[i] = make([]int, 10)
	}
	//对s中给定的字符串进行切分 转换成切片 对s中指定(已提供出)的位置置为 1
	str := strings.Split(s, " ") //str = ["1A" "2B" "1C"]
	for _, v := range str {
		i, _ := strconv.Atoi(v[:len(v)-1])
		p[i-1][v[len(v)-1]-'A'] = 1
	}
	res := 0
	//遍历整个位置p数组 由于我们只要计数,按照题意,可以得知
	//要容纳下四口之家 则情况有三种 1234 3456 5678 这三种
	//故j的值从 1~5
	for i := 0; i < n; i++ {
		for j := 1; j < 6; j++ {
			if j == 1 || j == 3 || j == 5 {
				if p[i][j] == 0 && p[i][j+1] == 0 && p[i][j+2] == 0 && p[i][j+3] == 0 {
					res++
					p[i][j], p[i][j+1], p[i][j+2], p[i][j+3] = 1, 1, 1, 1
				}
			}
		}
	}
	return res
}

func main() {
	Solution(2, "1A 2B 1C")
}
