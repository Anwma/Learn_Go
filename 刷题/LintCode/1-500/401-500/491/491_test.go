package _91

import (
	"testing"
)

//func IsPalindrome(num int) bool {
//	// write your code here
//	str := strconv.Itoa(num)
//	l,r := 0,len(str)-1
//	for l < r{
//		if str[l] != str[r]{
//			return false
//		}
//		l++
//		r--
//	}
//	return true
//}

func IsPalindrome (num int) bool {
	// write your code here
	tmp := num
	if tmp < 0 {
		return false
	}
	m := 0
	for tmp != 0 {
		m = m * 10 + tmp % 10
		tmp /= 10
	}
	return m == num
}

func BenchmarkIsPalindrome(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(12345678987654321)
	}
	b.StopTimer()
}