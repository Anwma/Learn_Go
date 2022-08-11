package _331

import "testing"

func EnglishSoftware(score []int, ask []int) []int {
	// write your code here
	sCount := len(score)
	r := make([]int, len(ask))
	for i, v := range ask {
		c := 0
		for _, s := range score {
			if s <= score[v-1] {
				c++
			}
		}
		r[i] = (c - 1) * 100 / sCount
	}
	return r
}

func BenchmarkEnglishSoftware(b *testing.B) {
	score := []int{100, 98, 87}
	ask := []int{1, 2, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EnglishSoftware(score, ask)
	}
	b.StopTimer()
}
