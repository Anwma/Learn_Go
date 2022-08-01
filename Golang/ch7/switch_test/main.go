package main

import "fmt"

func main() {
	score := 90
	grade := "A"
	/*
		90+   A
		80~90 B
		70~79 C
		60~69 D
		60-   E
	*/
	switch {
	case score >= 90:
		grade = "A"
		fmt.Println("优秀")
	case score >= 80 && score <= 89:
		grade = "B"
	case score >= 70 && score <= 79:
		grade = "C"
	case score >= 60 && score <= 69:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Println(grade)
	//不同的case表达式使用逗号分隔
	var a = "daddy"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
