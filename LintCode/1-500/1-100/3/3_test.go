package main

import (
	"strconv"
	"testing"
)

func DigitCounts(k int, n int) int {
	// write your code here
	//10000         45          25957320 ns/op
	//1000		  4318            268662 ns/op
	//100	    197233              5646 ns/op
	//10     3462813               346.1 ns/op
	//str := ""
	//for i := 0; i <= n; i++ {
	//	str += strconv.Itoa(i)
	//}
	//count := 0
	//for _, v := range str {
	//	//string 的本质是 []byte,取出来的是byte值
	//	//而k是int值,byte转为int值则是相应的ascii值。
	//	if k == int(v-'0') {
	//		count++
	//	}
	//}
	//return count

	//10000         757           1532715 ns/op
	//1000		   8143            148761 ns/op
	//100	      95143             12868 ns/op
	//10         968194              1286 ns/op
	//count := 0
	//strK := fmt.Sprintf("%d", k)//转成string
	//for i := 0; i <= n; i++ {
	//	//利用Split包切分出k的string,并计数。
	//	count += len(strings.Split(fmt.Sprintf("%d", i), strK)) - 1
	//}
	//return count

	//10000         31080             37756 ns/op
	//1000		   405276              3011 ns/op
	//100	    5044884               234.2 ns/op
	//10      63050377                19.45 ns/op
	//num := 0
	//if k == 0 {
	//	num++
	//}
	//for i := 0; i <= n; i++ {
	//	for m := i; m != 0; {
	//		j := m % 10
	//		m = m / 10
	//		if j == k {
	//			num++
	//		}
	//	}
	//}
	//return num

	//10000         9744246               122.3 ns/op
	//1000		   11785885               110.1 ns/op
	//100	      13299256                94.48 ns/op
	//10          19974631                59.75 ns/op
	str := strconv.Itoa(n)

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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DigitCounts(1, 10)
	}
	b.StopTimer()
}
