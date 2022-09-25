package main

import "testing"

func UniquePathsWithObstacles(a [][]int) int {
	// write your code here
	//cornet case
	if a == nil || len(a) == 0 || len(a[0]) == 0 {
		return 0
	}
	//初始化二维切片
	m, n := len(a), len(a[0])
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	//遍历整个二维切片
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//遇到障碍物 res数组(f[])置0
			if a[i][j] == 1 {
				//obstacle
				f[i][j] = 0
				//跳出本次循环
				continue
			}
			//当前点为初始点时，方案数为一
			if i == 0 && j == 0 {
				f[i][j] = 1
				continue
			}
			f[i][j] = 0
			if i > 0 {
				f[i][j] += f[i-1][j]
			}
			if j > 0 {
				f[i][j] += f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}

func BenchmarkUniquePathsWithObstacles(b *testing.B) {
	a := [][]int{{0, 0}, {0, 0}, {0, 0}, {1, 0}, {0, 0}}
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstacles(a)
	}
}
