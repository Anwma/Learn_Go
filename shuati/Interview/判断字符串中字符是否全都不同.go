package main

import "strings"

/*
请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允
许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都
不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓
度⼩于等于【3000】
*/
func main() {

}

func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		//Count 查找 s 中是否有重复的元素 有重复的元素则代表其count > 1 false
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		if v > 127 {
			return false
		}
		//Index 代表查找 s 中 某个字符 v 的值是否与其本身的 索引值不一样, 不一样返回false
		if strings.Index(s, string(v)) != k {
			return false
		}
	}
	return true
}
