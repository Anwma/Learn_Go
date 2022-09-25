package _69

import (
	"math"
	"testing"
)

func CoinChange(coins []int, amount int) int {
	// write your code here
	//确定状态 -子问题 -最后一步
	//f[i]代表i块钱时能换取的最少的硬币数量
	//转移方程：f[i] = min(f[i-1] + coins[?])
	f := make([]int, amount+1)
	n := len(coins)
	//初始条件和边界情况
	//f[0] = 0 f[0] ~ f[amount]
	//计算顺序 从小到大
	f[0] = 0
	for i := 1; i <= amount; i++ {
		//对f[i]中的值给定一个max值
		f[i] = math.MaxInt
		for j := 0; j < n; j++ {
			if i >= coins[j] && f[i-coins[j]] != math.MaxInt {
				//f[i-coins[j]] + 1
				f[i] = min(f[i], f[i-coins[j]]+1)
			}
		}
	}
	if f[amount] == math.MaxInt {
		return -1
	}
	return f[amount]

}
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	num := 27
	for i := 0; i < b.N; i++ {
		CoinChange(coins, num)
	}
}
