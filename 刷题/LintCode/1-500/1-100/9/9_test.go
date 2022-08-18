package main

import (
	"strconv"
	"testing"
)

func FizzBuzz(n int) []string {
	// write your code here
	//BenchmarkFizzBuzz-8      6568791               182.9 ns/op           240 B/op          1 allocs/op
	if n == 0 {
		return []string{}
	}
	str := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			str[i-1] = "fizz buzz"
		} else if i%5 == 0 {
			str[i-1] = "buzz"
		} else if i%3 == 0 {
			str[i-1] = "fizz"
		} else {
			str[i-1] = strconv.Itoa(i)
		}
	}
	return str
}

//func FizzBuzz(n int) []string {
//	if n == 0 {
//		return []string{}
//	}
//	//BenchmarkFizzBuzz-8      5937321               185.9 ns/op           240 B/op          1 allocs/op
//	str := make([]string, n)
//	for i := 1; i <= n; i++ {
//		switch {
//		case i%15 == 0:
//			str[i-1] = "fizz buzz"
//		case i%5 == 0:
//			str[i-1] = "buzz"
//		case i%3 == 0:
//			str[i-1] = "fizz"
//		default:
//			str[i-1] = strconv.Itoa(i)
//		}
//	}
//	return str
//}

//func FizzBuzz(n int) []string {
//	// write your code here
//	//BenchmarkFizzBuzz-8      6278198               201.8 ns/op           368 B/op          2 allocs/op
//	//提前用一个数组计数好 用1 2 3标记
//	flag := make([]int, n)
//	//将3的倍数的位置的值标记为1,index从0开始,所以i是2
//	for i := 2; i < n; i += 3 {
//		flag[i]++
//	}
//	//将5的倍数的位置的值标记为2,index从0开始,所以i是2
//	for i := 4; i < n; i += 5 {
//		flag[i] += 2
//	}
//
//	ans := make([]string, n)
//	for i := range ans {
//		switch flag[i] {
//		case 0:
//			ans[i] = strconv.Itoa(i + 1)
//		case 1:
//			ans[i] = "fizz"
//		case 2:
//			ans[i] = "buzz"
//		case 3:
//			ans[i] = "fizz buzz"
//		}
//	}
//	return ans
//}

func BenchmarkFizzBuzz(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FizzBuzz(15)
	}
	b.StopTimer()
}
