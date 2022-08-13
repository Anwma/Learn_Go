package main

import "fmt"

func main() {
	/*

		for expression1; expression2; expression3 {
			//...
		}

	*/

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}

	for index := 10; index > 0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}
	// break打印出来10、9、8、7、6
	// continue打印出来10、9、8、7、6、4、3、2、1
	map1 := make(map[int]int)
	for k, v := range map1 {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
	map2 := make(map[int]int)
	for _, v := range map2 {
		fmt.Println("map's val:", v)
	}
}
