package _910

import (
	"sort"
	"testing"
)

func FindNumber(array []int) int {
	length := len(array)
	m := make(map[int]int, length)
	for i := 0; i < length; i++ {
		_, ok := m[array[i]]
		if !ok {
			m[array[i]] = 1
		} else {
			m[array[i]]++
		}
	}

	keys := []int{}
	for key := range m {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	maxIndex := keys[0]
	for _, key := range keys {
		if m[key] > m[maxIndex] {
			maxIndex = key
		}
	}
	return maxIndex
}

//func FindNumber(array []int) int {
//	//感觉可以用多个变量 DP写 但是太菜 写不出来...
//	return 0
//}

func BenchmarkFindNumber(b *testing.B) {
	arr := []int{1, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindNumber(arr)
	}
	b.StopTimer()
}
