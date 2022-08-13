package main

import "fmt"

func main() {
	//1-3
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//7-9
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	//loop 相当于 while(true){break}
	for {
		fmt.Println("loop")
		break
	}
	//打印奇数(0-5之间)
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
