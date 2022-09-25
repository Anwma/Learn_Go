package main

/*
请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
符串(可以使⽤单个过程变量)。
给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
5000。
*/

// string 为返回的字符 bool 为其长度是否小于5000
func reverseString(s string) (string, bool) {
	str := []rune(s)
	l := len(str)
	//返回原始字符串 和 false 值
	if l > 5000 {
		return s, false
	}
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str), true
}
