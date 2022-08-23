package main

import (
	"math"
	"strconv"
	"testing"
)

//func ReverseInteger(n int) int {
//	var result int64
//
//	for ; n != 0; n /= 10 {
//		result = result*10 + int64(n%10)
//		if result > math.MaxInt32 || result < math.MinInt32 {
//			return 0
//		}
//	}
//	return int(result)
//}


func ReverseInteger(n int) int {
	// write your code here
	str := strconv.Itoa(n)
	//由于字符串无法进行修改 故将其转化为[]byte切片
	strSlice := []byte(str)
	length := len(strSlice) - 1
	//负数
	var temp []byte
	if strSlice[0] == '-' {
		temp = reverse(strSlice, 1, length)
	} else {
		temp = reverse(strSlice, 0, length)
	}
	//正数
	res, _ := strconv.Atoi(string(temp))
	if res > math.MaxInt32 {
		return 0
	}
	return res
}

func reverse(slice []byte, left, right int) []byte {
	for left <= right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
	return slice
}

func BenchmarkReverseInteger(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseInteger(-123456789)
	}
	b.StopTimer()
}
