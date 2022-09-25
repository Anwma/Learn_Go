package main

import "strings"

/*
给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string
s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的
⻓度都⼩于等于5000。
*/
func isRegroup(s1, s2 string) bool {
	sL1 := len([]rune(s1))
	sL2 := len([]rune(s2))
	//长度超过5000 以及 二者长度不一
	if sL1 > 5000 || sL2 > 5000 || sL1 != sL2 {
		return false
	}
	for _, v := range s1 {
		//计算s1 中某些字符的数量是否和 s2 中对应的字符数量不一样 如果不一样则返回 false
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}
