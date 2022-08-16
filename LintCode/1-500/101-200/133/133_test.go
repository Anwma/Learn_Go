package main

import "testing"

func LongestWords(str []string) []string {
	//给一个词典，找出其中所有最长的单词。
	if str == nil || len(str) <= 1 {
		return str
	}

	res := []string{str[0]}
	for i := 1; i < len(str); i++ {
		//当前的字符串长度比原先的字符串长度长
		if len(str[i]) > len(res[0]) {
			//先清空原先表内的内容 如何用slice做到呢?
			res = res[:0]
			res = append(res, str[i])
		} else if len(str[i]) == len(res[0]) {
			//长度相等，则添加到切片当中
			res = append(res, str[i])
		} else {
			//长度小于
			continue
		}
	}
	return res
}

func BenchmarkLongestWords(b *testing.B) {
	str := []string{"dog", "google", "facebook", "internationalization", "blabla"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestWords(str)
	}
	b.StopTimer()
}

