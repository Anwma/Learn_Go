package main

import "testing"

// DP
func CanJump(a []int) bool {
	// write your code here
	//1.确定状态 最后一步 子问题
	//共 n 块石头, 如果青蛙能跳到第n块石头上,那么往前一步,青蛙能跳到第i块石头上
	//0<=i<n, 那么青蛙一定能跳到第i块石头,且当前第i块石头的值 + i >= n

	//2.转移方程 bool : f[n] = f[i] & i + f[i] >= n
	//3.初始条件及边界情况 f[0] = 0
	//4.计算顺序 从小到大
	n := len(a)
	f := make([]bool, n) //开一个bool切片，存储res值
	f[0] = true
	for j := 1; j < n; j++ {
		f[j] = false
		for i := 0; i < j; i++ {
			if f[i] && i+a[i] >= j {
				f[j] = true
				break
			}
		}
	}
	return f[n-1]
}

// 贪心
func CanJump1(a []int) bool {
	far := 0 //最远
	n := len(a)
	for i := 0; i < n; i++ {
		if far < i {
			return false
		}
		far = max(i+a[i], far) //直接利用far来继承上一步到达的最高值 O(n)
		//i = far - 1 如果可以确定数组中的元素均大于0 则可以快速到达终点处
		//但是因为要判断是否可达，也就是说a[]中的数据不能确保均大于0 所以此处未进行加速
		if far >= n-1 {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func BenchmarkCanJump(b *testing.B) {
	a := []int{1, 0}
	for i := 0; i < b.N; i++ {
		CanJump1(a)
		//0		22.65 ns/op
		//1		2.381 ns/op
	}
}
