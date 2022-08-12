package _535

import (
	"testing"
)

func ToLowerCase(str string) string {
	// Write your code here

	//return strings.ToLower(str)
	//BenchmarkToLowerCase-8          2648 2348                38.98 ns/op

	s := []byte(str)
	for i := 0; i < len(s); i++ {
		if s[i] <= 'Z' && s[i] >= 'A' {
			s[i] += 32
		}
	}
	return string(s)
	//BenchmarkToLowerCase-8          1 0000 0000               11.69 ns/op

}

func BenchmarkToLowerCase(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToLowerCase("Hello")
	}
	b.StopTimer()
}
