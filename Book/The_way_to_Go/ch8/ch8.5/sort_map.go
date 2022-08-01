// the telephone alphabet:
package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	//map 是无序的
	fmt.Println(barVal)
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
	//创建一个keys-string slice
	keys := make([]string, len(barVal))
	i := 0
	//将barVal 中的key放到keys[i]-value-slice 中
	for k := range barVal {
		keys[i] = k
		i++
	}
	//对keys 排序
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	//for-range keys keys中的key=index,value = k 也就是所包含的值
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v\n", k, barVal[k])
	}
}
