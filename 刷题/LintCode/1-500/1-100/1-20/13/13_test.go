package main

import (
	"strings"
	"testing"
)

func StrStr(source string, target string) int {
	// Write your code here
	return strings.Index(source, target)
}

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrStr("1234", "12")
	}
}
