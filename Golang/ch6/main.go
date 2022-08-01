package main

import (
	"fmt"
	"strings"
)

func main() {
	//字符串基本操作
	//1. 求解字符串的长度
	//返回的是字节的长度
	//涉及到中文问题就产生了变化
	//unicode 字符集, 存储的时候需要编码 utf8编码规则 utf8是一个动态的编码规则
	//utf8编码,还能够用一个字节表示英文
	//var name = "wozen:楚心云" //转义符
	//fmt.Println(len(name))
	//类型转化 转换为rune数组
	//name_arr := []int32(name)
	//fmt.Println(len(name_arr))
	//date := "2022\\07\\14"
	//date := `2022\07\14`
	//fmt.Println(date)

	//2.是否包含某个子串
	var name = "wozenw:楚心云" //转义符
	fmt.Println(strings.Contains(name, "wozen"))
	fmt.Println(strings.Index(name, "楚心云"))

	//3.统计出现的次数
	fmt.Println(strings.Count(name, "w"))

	//4.前缀和后缀
	fmt.Println(strings.HasPrefix(name, "w"))
	fmt.Println(strings.HasSuffix(name, "云"))

	//5.大小写转换
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("WOZEN"))

	//6.字符串的比较
	fmt.Println(strings.Compare("aa", "ab"))
	fmt.Println(strings.Compare("b", "a"))
	fmt.Println(strings.Compare("b", "b"))
	//字符串的比较就是ascii的比较 返回-1(第一个字符小于第二个字符) 0 1(大于)

	//7.去掉空格和指定字符串
	fmt.Println(strings.TrimSpace(" wozen "))   //去掉首尾空格
	fmt.Println(strings.TrimLeft("hello", "h")) //去掉左端指定字符串
	fmt.Println(strings.Trim("bobby", "b"))     //去掉首尾两端指定字符串

	//8.split方法
	fmt.Println(strings.Split("hello world", " ")) //分割成数组

	//9.合并 join方法将字符串数组连接起来
	arrs := strings.Split("hello world", " ")
	fmt.Println(strings.Join(arrs, ","))

	//10.字符串替换 s:源字符串 old:s中要替换的字符串 new:替换后的新字符串 n:替换的次数
	fmt.Println(strings.Replace("hello world wor", "wor", "111", 2))
}
