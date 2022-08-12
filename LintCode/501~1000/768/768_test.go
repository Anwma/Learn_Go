package _68

import "testing"

func CalcYangHuisTriangle(n int) [][]int {
	// write your code here
	//tmp := []int{1}
	//res := [][]int{}
	//for i := 0; i < n; i++ {
	//	res = append(res, tmp)
	//	//每行的元素个数与当前行号相等
	//	//下一行的第i个元素为上一行的第j-1个元素与第i个元素之和
	//	for j := 1; j < i; j++ {
	//		res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
	//	}
	//	//其他的元素赋初值
	//	if i > 0 {
	//		res[i] = append(res[i], 1)
	//	}
	//}
	//return res

	triangle := make([][]int, n)
	for i, _ := range triangle {
		triangle[i] = make([]int, i+1)
		// 外层
		triangle[i][0] = 1
		triangle[i][i] = 1
		// 内层
		if i >= 2 {
			for j := 1; j < len(triangle[i])-1; j++ {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}
	return triangle
	//185.7 ns/op

	//ans := make([][]int,n)
	//for i,_ := range ans {
	//	ans[i] = make([]int,i + 1)
	//	ans[i][0] = 1
	//	ans[i][i] = 1
	//}
	////dp
	//for i := 2; i < n; i++ {
	//	for j := 1; j < len(ans[i]) - 1; j++ {
	//		ans[i][j] = ans[i - 1][j - 1] + ans[i - 1][j]
	//	}
	//}
	//return ans
}

func BenchmarkCalcYangHuisTriangle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CalcYangHuisTriangle(4)
	}
	b.StopTimer()
}
