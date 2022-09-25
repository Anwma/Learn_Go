package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	grid := [][]int{{1, 3, 1, 4, 2, 5}, {1, 5, 9, 3, 1, 4}, {4, 2, 1, 3, 1, 7}}
	MinPathSum1(grid)
}
func MinPathSum1(grid [][]int) int {
	// write your code here
	m, n := len(grid), len(grid[0])
	if grid == nil || m == 0 || n == 0 {
		return 0
	}
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}
	//pi存放当前是往左走还是往右走
	pi := make([][]int, m)
	for i := 0; i < m; i++ {
		pi[i] = make([]int, n)
	}
	//往左走：pi[i][j] = 0
	//往右走：pi[i][j] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[i][j]
				continue
			}
			f[i][j] = math.MaxInt
			if i > 0 {
				f[i][j] = min1(f[i][j], f[i-1][j]+grid[i][j])
				if f[i][j] == f[i-1][j]+grid[i][j] {
					pi[i][j] = 0
				}
			}
			if j > 0 {
				f[i][j] = min1(f[i][j], f[i][j-1]+grid[i][j])
				if f[i][j] == f[i][j-1]+grid[i][j] {
					pi[i][j] = 1
				}
			}
		}
	}
	//通过pi[i][j]我们可以区分出到底f[i][j]来源于上一步的往左还是往右
	//创建路径数组
	path := make([][]int, m+n-1)
	for i := 0; i < m+n-1; i++ {
		path[i] = make([]int, 2)
	}
	//
	i, j := m-1, n-1
	for p := m + n - 2; p >= 0; p-- {
		path[p][0] = i
		path[p][1] = j
		if p == 0 {
			break
		}
		if pi[i][j] == 0 {
			i--
		} else {
			j--
		}
	}
	//打印路径
	for p := 0; p < m+n-1; p++ {
		fmt.Println("(" + strconv.Itoa(path[p][0]) + ", " + strconv.Itoa(path[p][1]) + "):" + strconv.Itoa(grid[path[p][0]][path[p][1]]))
	}
	return f[m-1][n-1]
}

func min1(x, y int) int {
	if x > y {
		return y
	}
	return x
}
