package main

import (
	"sort"
	"testing"
)

func FindContentChildren(g []int, s []int) int {
	// Write your code here
	//1.对数组 g s 进行排序,保证数组的有序性
	sort.Ints(g)
	sort.Ints(s)
	//2.初始化孩子和饼干数
	//child cookie 还代表着二者的index值(复用,避免多分配变量空间)
	child, cookie := 0, 0
	for child < len(g) && cookie < len(s) {
		//若能被喂饱 则二者index均往后走一位
		if g[child] <= s[cookie] {
			child++
			cookie++
		} else {
			//若当前饼干不足以喂饱最低饥饿度的孩子，则饼干索引值++
			cookie++
		}
	}
	return child
}

func BenchmarkFindContentChildren(b *testing.B) {
	g := []int{1, 2, 3}
	s := []int{1, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindContentChildren(g, s)
	}
	b.StopTimer()
}
