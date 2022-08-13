package _45

import "testing"

func LowercaseToUppercase(character byte) byte {
	// write your code here
	return character-32
}

func BenchmarkLowercaseToUppercase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LowercaseToUppercase('a')
	}
	b.StopTimer()
}