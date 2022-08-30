package _69

import (
	"testing"
)

func BenchmarkCoinChange(b *testing.B) {
	coins := []int{1, 2, 5}
	num := 27
	for i := 0; i < b.N; i++ {
		CoinChange(coins, num)
	}
}
func CoinChange(coins []int, amount int) int {
	arr := make([]int, amount+1)
	n := len(coins)
	arr[0] = 0
	for i := 1; i <= amount; i++ {
		arr[i] = 2<<31 - 1
		//last coin
		//arr[i] = min{arr[i-coins[0]]+1, ... , f[i-coins[n-1]] + 1}
		for j := 0; j < n; j++ {
			//i 代表当前的钱数 当前钱数要 >= 硬币数组中的值
			//且
			if i >= coins[j] {
				arr[i] = min(arr[i-coins[j]]+1, arr[i])
			}
		}
	}
	if arr[amount] == 2<<31-1 {
		arr[amount] = -1
	}
	return arr[amount]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
