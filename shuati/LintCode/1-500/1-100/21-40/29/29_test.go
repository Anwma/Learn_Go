package main

import "testing"

func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	memo := make([][]int, len(s1)+1)
	for i := range memo {
		memo[i] = make([]int, len(s2)+1)
	}
	return helper(s1, s2, s3, 0, 0, 0, memo)
}

func helper(s1, s2, s3 string, i1, i2, i3 int, memo [][]int) bool {
	if i3 == len(s3) {
		return true
	}

	if memo[i1][i2] != 0 {
		return memo[i1][i2] == 1
	}

	if i1 < len(s1) && s1[i1] == s3[i3] {
		if helper(s1, s2, s3, i1+1, i2, i3+1, memo) {
			memo[i1][i2] = 1
			return true
		}
	}
	if i2 < len(s2) && s2[i2] == s3[i3] {
		if helper(s1, s2, s3, i1, i2+1, i3+1, memo) {
			memo[i1][i2] = 1
			return true
		}
	}
	memo[i1][i2] = -1
	return false
}

func BenchmarkIsInterleave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsInterleave("aabcc", "dbbca", "aadbbcbcac")
	}
}
