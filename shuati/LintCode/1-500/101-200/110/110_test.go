package main

import "math"

func MinPathSum(grid [][]int) int {
	// write your code here
	m, n := len(grid), len(grid[0])
	if grid == nil || m == 0 || n == 0 {
		return 0
	}
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[i][j]
				continue
			}
			f[i][j] = math.MaxInt
			if i > 0 {
				f[i][j] = min(f[i][j], f[i-1][j]+grid[i][j])
			}
			if j > 0 {
				f[i][j] = min(f[i][j], f[i][j-1]+grid[i][j])
			}
		}
	}
	return f[m-1][n-1]
}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
