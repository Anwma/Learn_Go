package _138

import "testing"

func CanPlaceFlowers(flowerbed []int, n int) bool {
	// Write your code here
	length := len(flowerbed)
	for i, v := range flowerbed {
		//首尾部 当前值为0,且首部下一个index==0,尾部上一个index==0。可进行放置
		//中部 当前值为0 其前面一个和后面一个index均为0 方可放置
		if v == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			n--
			//并将当前值置为1
			flowerbed[i] = 1
			if n <= 0 {
				return true
			}
		}
	}
	//n 为 0的情况
	return n == 0
}

func BenchmarkCanPlaceFlowers(b *testing.B) {
	flowerbed := []int{1, 0, 0, 0, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CanPlaceFlowers(flowerbed, 1)
	}
	b.StopTimer()
}
