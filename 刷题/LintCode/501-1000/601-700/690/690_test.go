package _90

import (
	"testing"
)

//func Factorial(n int) string {
//	// write your code here
//	//n过大 不行
//	if n <= 1 {
//		return "1"
//	}
//	var res int64 = 1
//	for i := 2; i <= n; i++ {
//		res *= int64(i)
//	}
//	return strconv.FormatInt(res, 10)
//}

import "fmt"

func Factorial (n int) string {
	// write your code here
	res := []int{1}
	for i := 2; i <= n; i++ {
		var carry int
		for j := 0; j < len(res); j++ {
			tmp := res[j] * i + carry
			res[j] = tmp % 10
			carry = tmp / 10
		}
		for carry > 0 {
			res = append(res, carry % 10)
			carry = carry/ 10
		}
	}
	var str string
	for i := len(res)-1; i>= 0; i-- {
		str += fmt.Sprintf("%d", res[i])
	}
	return str
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(30)
	}
}
