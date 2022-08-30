package main

import (
	"fmt"
	"testing"
)

func BenchmarkIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := IsUnique("a")
		fmt.Print(res)
	}
}
func IsUnique(str string) bool {
	// write your code here
	m := make(map[string]bool)
	for i := 0; i< len(str); i++{
		val := string(str[i])
		if m[val]{
			return false
		}
		m[val] = true
	}
	return true
}