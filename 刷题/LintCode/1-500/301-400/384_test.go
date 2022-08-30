package main

import "testing"

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb"
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring(s)
	}
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func LengthOfLongestSubstring(s string) int {
	// write your code here
	//corner case
	lens := len(s)
	if lens <= 1 {
		return lens
	}
	//normal
	//1.创建一个map 保存 字符及其出现的索引位置
	m := make(map[string]int)
	start, res := 0, 0
	//遍历整个字符串 s
	for i := 0; i < lens; i++ {
		tmp := string(s[i])
		if _, ok := m[tmp]; ok {
			//存在 则做指针往右移动到该重复字符出现的index位置上 存放
			start = max(start, m[tmp]+1)
			res = max(res, i-start+1)
			m[tmp] = i
		} else {
			//不存在 存放
			m[tmp] = i
			res = max(res, i-start+1)
		}
	}
	return res
}
