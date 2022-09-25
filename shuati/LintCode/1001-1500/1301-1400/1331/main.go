package main

import "fmt"

func main() {
	score := []int{100, 101, 102, 103, 104, 98, 87}
	ask := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print(EnglishSoftware(score, ask))
}

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
