package main

import "testing"

func RotateString(s []byte, offset int) {
	// write your code here
	if offset == 0 || len(s) == 0 {
		return
	}
	offset %= len(s)
	reverse(s, 0, len(s)-offset-1)
	reverse(s, len(s)-offset, len(s)-1)
	reverse(s, 0, len(s)-1)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func BenchmarkRotateString(b *testing.B) {
	s := []byte{'a', 'b', 'c'}
	for i := 0; i < b.N; i++ {
		RotateString(s, 2)
		//0		13.80 ns/op
	}
}
