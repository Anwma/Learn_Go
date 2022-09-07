package main

import (
	"fmt"
	"testing"
)

func MinimalOperation(words []string) []int {
	// Write your code here.
	//1.创建一个[]int
	res := make([]int, len(words))
	//2.相邻的字符间有重复的
	for i, word := range words {
		ans := 0
		x := 1
		for j := 0; j < len(word)-1; j++ {
			if word[j] == word[j+1] {
				x++
			}
			if word[j] != word[j+1] {
				ans += x / 2
				x = 1
			}
		}
		ans += x / 2
		res[i] = ans
	}
	return res
}
func BenchmarkMinimalOperation(b *testing.B) {
	arr := []string{"ab", "aab", "abb", "abab", "abaaaaba"}
	for i := 0; i < b.N; i++ {
		res := MinimalOperation(arr)
		fmt.Print(res)
	}
}

/*
    res := make([]int, len(words))
	for index, word := range words {
		m := make(map[byte]bool)
		for _, item := range word {
			if _, ok := m[byte(item)]; ok {
				res[index]++
			} else {
				m[byte(item)] = true
			}
		}
	}
	return res
*/
