package main

import "math"

func MinCost(costs [][]int) int {
	// write your code here
	n := len(costs)
	if n == 0 {
		return 0
	}
	f := make([][]int, n+1)
	//序列型
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, 3)
	}
	//初始状态
	f[0][0], f[0][1], f[0][2] = 0, 0, 0
	//前 i 栋 房子
	for i := 1; i <= n; i++ {
		for j := 0; j < 3; j++ {
			f[i][j] = math.MaxInt
			for k := 0; k < 3; k++ {
				//不能撞色
				if j == k {
					continue
				}
				f[i][j] = min(f[i][j], f[i-1][k]+costs[i-1][j])
			}
		}
	}
	return min(f[n][0], min(f[n][1], f[n][2]))
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
