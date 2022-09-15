package main

import "math"

func MinPathSum(grid [][]int) int {
	// write your code here
	m, n := len(grid), len(grid[0])
	if grid == nil || m == 0 || n == 0 {
		return 0
	}
	//初始化滚动数组
	f := make([][]int, 2)
	for i := 0; i < 2; i++ {
		f[i] = make([]int, n)
	}
	//开两个变量
	old, now := 0, 0
	//old f[old][...] is holding f[i-1][...]
	//now f[now][...] is holding f[i][...]
	for i := 0; i < m; i++ {
		old = now
		now = 1 - old
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[now][j] = grid[i][j]
				continue
			}
			f[i][j] = math.MaxInt
			if i > 0 {
				f[now][j] = min(f[now][j], f[old][j]+grid[i][j])
			}
			if j > 0 {
				f[now][j] = min(f[now][j], f[now][j-1]+grid[i][j])
			}
		}
	}
	return f[now][n-1]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
