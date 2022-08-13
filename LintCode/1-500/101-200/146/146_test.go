package _46

import (
	"testing"
)

func LowercaseToUppercase2(letters string) string {
	// write your code here
	//return strings.ToUpper(letters)
	strBytes := []byte(letters)
	for i := range strBytes {
		if strBytes[i] >= 'a' && letters[i] <= 'z' {
			strBytes[i] -= 32
		}
	}
	return string(strBytes)
}

func BenchmarkLowercaseToUppercase2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LowercaseToUppercase2("abc")
	}
	b.StopTimer()
}
