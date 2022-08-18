package _20

import "testing"

func GetAnswer(num int) int {
	// write your code here.
	count := 0
	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = 3*num + 1
		}
		count++
	}
	return count
}

func BenchmarkGetAnswer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetAnswer(100)
	}
	b.StopTimer()
}
