package main

import (
	"strconv"
	"strings"
	"testing"
)

func DigitCounts(k int, n int) int {
	// write your code here
	count := 0
	strK := strconv.Itoa(k)
	for i := 0; i <= n; i++ {
		count += len(strings.Split(strconv.Itoa(i), strK)) - 1
	}
	return count
}

func DigitCounts1(k int, n int) int {
	// write your code here
	str := strconv.Itoa(n)
	//n的取值范围为0~10^5
	maxBit := 6
	tenPow := make([]int, maxBit)
	tenPow[0] = 1
	for i := 1; i < maxBit; i++ {
		tenPow[i] = tenPow[i-1] * 10
	}
	res := 0
	for i := 0; i < len(str); i++ {
		curNum := int(str[i] - '0')
		count := 0
		// i位置前的可能性
		if i > 0 {
			if k != 0 {
				count = n / tenPow[len(str)-i] * tenPow[len(str)-1-i]
			} else {
				count = (n/tenPow[len(str)-i] - 1) * tenPow[len(str)-1-i]
			}
		}
		// i位置后面的可能性
		if i > 0 || k != 0 {
			if curNum > k {
				count += tenPow[len(str)-i-1]
			} else if curNum == k {
				count += n%tenPow[len(str)-i-1] + 1
			}
		}
		res += count
	}
	if k == 0 {
		res++
	}
	return res
}

func BenchmarkDigitCounts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitCounts1(1, 11)
	}
}
