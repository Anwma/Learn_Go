package main

import (
	"strings"
	"unicode"
)

/*
请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩
写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。
*/
func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	//n:-1 表示 不限制替换数量
	return strings.Replace(s, " ", "%20", -1), true
}
