package _5

import "testing"

//func PrintX(n int) []string {
//	// write your code here
//	//1.创建一个长度为 n 的 []string
//	res := make([]string, n)
//	//2.针对每一个[]里面,对其进行操作
//	for i := 0; i < n; i++ {
//		//3.每一个切片中,开辟大小为n的byte[]
//		item := make([]byte, n)
//		//4.每个[]byte中的item初始化为 ' '
//		for j := range item {
//			item[j] = ' '
//		}
//		//5.只需对当前的行切片[]byte的首尾两个index赋值为 'X' 即可
//		item[i] = 'X'
//		item[n-i-1] = 'X'
//		//6.对 []string 赋相应的 []byte
//		res[i] = string(item)
//	}
//	//7.返回结果
//	return res
//}

func PrintX(n int) []string {
	// write your code here
	res := []string{}
	for i := 0; i < n; i ++ {
		str := ""
		for j := 0; j < n; j ++ {
			if valid(i, j, n) {
				str += "X"
			} else {
				str += " "
			}
		}
		res = append(res, str)
	}

	return res
}

func valid(i,j,k int) bool {
	if i == j {
		return true
	}
	if i + j + 1 == k {
		return true
	}
	return false
}

func BenchmarkPrintX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintX(15)
	}
	b.StopTimer()
}