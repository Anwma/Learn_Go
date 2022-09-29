package main

import "fmt"

type WeightNode struct {
	addr            string
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

type LoadBalance interface {
	Add(...string) error
	Get(string) (string, error)

	//后期服务发现补充
	Update()
}

func main() {
	for i := 1; i <= 100; i++ {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Print(i)
		fmt.Print(",")
	}
}
