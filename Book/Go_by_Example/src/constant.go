package main

import (
	"fmt"
	"math"
)

const s string = "constant"

//常量不允许修改

func main() {
	fmt.Println(s)

	const n = 500000000 //5*10^8

	const d = 3e20 / n //3*10^20
	fmt.Println(d)     //0.6*10^12 = 6 * 10^11

	fmt.Println(int64(d))
	//math.sin需要float64 n会自动确认类型
	fmt.Println(math.Sin(n))
}
