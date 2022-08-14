package main

func TrailingZeros(n int64) int64 {
	// write your code here
	ans := int64(0)
	for n > 0 {
		n /= 5
		ans += n
	}
	return ans
}