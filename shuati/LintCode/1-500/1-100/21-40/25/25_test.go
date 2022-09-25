package _5

import "testing"

func PrintX(n int) []string {
	// write your code here
	//1.创建一个长度为 n 的 []string
	res := make([]string, n)
	//2.针对每一个[]里面,对其进行操作
	for i := 0; i < n; i++ {
		//3.每一个切片中,开辟大小为n的byte[]
		item := make([]byte, n)
		//4.每个[]byte中的item初始化为 ' '
		for j := range item {
			item[j] = ' '
		}
		//5.只需对当前的行切片[]byte的首尾两个index赋值为 'X' 即可
		item[i] = 'X'     //首
		item[n-i-1] = 'X' //尾
		//6.对 []string 赋相应的 []byte
		res[i] = string(item)
	}
	//7.返回结果
	return res
}

func BenchmarkPrintX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintX(10)
	}
}
