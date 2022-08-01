package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	//longCalculation() 某个函数
	fb(35)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fb(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fb(n-1) + fb(n-2)
	}

}
