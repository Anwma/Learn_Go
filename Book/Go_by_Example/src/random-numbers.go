package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数0-100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	//生成随机float64位浮点数(0.0-1.0)
	fmt.Println(rand.Float64())
	//5.0-10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()
	//使用相同种子生成的随机数生成器，会生成相同的随机数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	//此处①
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	//此处②同①生成相同的两个随机数
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
