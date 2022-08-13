package _3

import (
	"strings"
	"testing"
)

func StrStr(source string, target string) int {
	//Write your code here
	return strings.Index(source, target)
}

func BenchmarkStrStr(b *testing.B) {
	source := "source"
	target := "target"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrStr(source, target)
	}
	b.StopTimer()
}
