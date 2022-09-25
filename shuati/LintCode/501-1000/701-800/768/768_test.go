package main

import "testing"

func CalcYangHuisTriangle(n int) [][]int {
	// write your code here
	//6223146               190.4 ns/op
	//初始化二维切片
	res := make([][]int, n)
	for i, _ := range res {
		//每行的元素个数为该行的行数 由于i从0开始索引 故+1
		res[i] = make([]int, i+1)
		//观察可得首尾两个元素为1
		res[i][0] = 1
		res[i][i] = 1
	}

	//从第3行开始 对应到index 就是 2
	for i := 2; i < n; i++ { //行
		//计算到当前行的倒数第二个元素
		for j := 1; j < len(res[i])-1; j++ { //列
			//当前元素为上一行的两个元素之和
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func BenchmarkCalcYangHuisTriangle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcYangHuisTriangle(4)
	}
	b.StopTimer()
}
