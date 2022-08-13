package echo6

import (
	"strings"
	"testing"
)

func StringJoin() string {
	s1 := []string{"hello", "world"}
	return strings.Join(s1, ",")
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}

func StringAdd() string {
	s1 := "hello"
	s2 := "world"
	return s1 + "," + s2
}

func BenchmarkStringAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringAdd()
	}
}
