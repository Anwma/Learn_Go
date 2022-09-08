package main

import "testing"

func IsAlphanumeric(c byte) bool {
	// write your code here
	if c <= 'Z' && c >= 'A' || c <= 'z' && c >= 'a' || c <= '9' && c >= '0' {
		return true
	}
	return false
}

func BenchmarkIsAlphanumeric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAlphanumeric('a')
	}
}
