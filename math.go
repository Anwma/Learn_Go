package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1; i < 3000; i++ {
		pos := 1 - math.Pow(149.0/150, float64(i))
		fmt.Print("第")
		fmt.Print(i)
		fmt.Print("次：")
		fmt.Println("综合中奖率为: ", pos*100, "%")
	}
}
