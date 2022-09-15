package main

import "testing"

func BenchmarkMaxKilledEnemies(b *testing.B) {
	test := [][]byte{}
	for i := 0; i < b.N; i++ {
		MaxKilledEnemies(test)
	}
}
func MaxKilledEnemies(grid [][]byte) int {
	// write your code here
	//corner case
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	//初始化 上下左右四个数组
	up := make([][]int, m)
	for i := 0; i < m; i++ {
		up[i] = make([]int, n)
	}
	down := make([][]int, m)
	for i := 0; i < m; i++ {
		down[i] = make([]int, n)
	}
	left := make([][]int, m)
	for i := 0; i < m; i++ {
		left[i] = make([]int, n)
	}
	right := make([][]int, m)
	for i := 0; i < m; i++ {
		right[i] = make([]int, n)
	}
	//up
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 'W' {
				if grid[i][j] == 'E' {
					up[i][j]++
				}
				if i > 0 {
					up[i][j] += up[i-1][j]
				}
			}
		}
	}
	//down
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if grid[i][j] != 'W' {
				if grid[i][j] == 'E' {
					down[i][j]++
				}
				if i < m-1 {
					down[i][j] += down[i+1][j]
				}
			}
		}
	}
	//left
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 'W' {
				if grid[i][j] == 'E' {
					left[i][j]++
				}
				if j > 0 {
					left[i][j] += left[i][j-1]
				}
			}
		}
	}
	//right
	for i := 0; i < m; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] != 'W' {
				if grid[i][j] == 'E' {
					right[i][j]++
				}
				if j < n-1 {
					right[i][j] += right[i][j+1]
				}
			}
		}
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' { //empty
				res = max(res, up[i][j]+down[i][j]+left[i][j]+right[i][j])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
