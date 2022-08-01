package main

import "fmt"

func main() {
	//创建切片nums len = 3
	nums := []int{2, 3, 4}
	sum := 0
	//对nums中的数据求和,存放在num中
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//取出num ==3时,对应的索引值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//创建key-value的map
	kvs := map[string]string{"a": "apple", "b": "banana", "p": "peach"}
	//随机取值 map无序
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		//字符串 输出索引值 0 1 和其对应的ascii码值
		fmt.Println(i, c)
		//Output:
		//0 103
		//1 111
	}
}
