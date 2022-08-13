package main

import "fmt"

func main() {
	s := printNumWith2(2.2, 3.2)
	fmt.Println(s)
	data := []byte{'a', 'b'}
	printBytes(data)
}

// 输出两位小数
func printNumWith2(a float64, b float64) string {
	s := fmt.Sprintf("%.2f,%.2f", a, b)
	return s
}

func printBytes(data []byte) string {
	for _, value := range data {
		fmt.Printf("%x ", value)
	}
	return ""
}
