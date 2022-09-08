package main

import "testing"

func NthUglyNumber(n int) int {
	uglies := make([]int, n)
	uglies[0] = 1
	pugly2, pugly3, pugly5 := 0, 0, 0
	i := 1
	for i < n {
		curUgly := min(min(uglies[pugly2]*2, uglies[pugly3]*3), uglies[pugly5]*5)
		uglies[i] = curUgly
		for uglies[pugly2]*2 <= curUgly {
			pugly2++
		}
		for uglies[pugly3]*3 <= curUgly {
			pugly3++
		}
		for uglies[pugly5]*5 <= curUgly {
			pugly5++
		}
		i++
	}
	return uglies[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func BenchmarkNthUglyNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NthUglyNumber(12)
	}
}
