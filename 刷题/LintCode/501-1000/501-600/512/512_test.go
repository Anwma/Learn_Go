package main

import "testing"

func NumDecodings(s string) int {
	// write your code here
	str := []byte(s)
	n := len(str)
	if n == 0 {
		return 0
	}
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		f[i] = 0
		//最后一位
		if str[i-1] != '0' {
			f[i] += f[i-1]
		}
		//最后两位
		if i >= 2 && (str[i-2] == '1' || (str[i-2] == '2' && s[i-1] <= '6')) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

func BenchmarkNumDecodings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NumDecodings("12345")
	}
}
