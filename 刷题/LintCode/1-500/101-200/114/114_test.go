package _14

import "testing"

func UniquePaths(m int, n int) int {
	// write your code here
	f := make([][]int, m)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n)
	}
	for i := 0; i < len(f); i++ {
		for j := 0; j < len(f[0]); j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}

func BenchmarkUniquePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UniquePaths(7, 5)
	}
}
