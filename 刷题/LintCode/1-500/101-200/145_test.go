package main

import "testing"

func LowercaseToUppercase(character byte) byte {
	// write your code here
	return character - 32
}

func BenchmarkLowercaseToUppercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LowercaseToUppercase('a')
	}
}
