package main

import (
	"fmt"
	"sort"
)

//构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；先打印所有的饮料，
//然后打印原名和翻译后的名字。接下来按照英文名排序后再打印出来。
func main() {
	juice := map[string]string{"tea": "茶", "apple-tea": "苹果茶", "coffee": "咖啡", "coke": "可乐"}
	//1. 先打印所有饮料
	fmt.Println(juice)
	//2. 打印原名和翻译后的名字
	for k, v := range juice {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}
	//sort
	str := make([]string, len(juice))
	i := 0
	for k := range juice {
		str[i] = k
		i++
	}
	fmt.Println("排序后：")
	sort.Strings(str)
	for _, k := range str {
		fmt.Printf("key: %v, value: %v\n", k, juice[k])
	}

}
