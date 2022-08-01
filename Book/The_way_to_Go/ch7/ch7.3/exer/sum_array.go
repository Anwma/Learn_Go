package main

import "fmt"

func main() {
	arr := make([]float32, 4)
	for i := 0; i < len(arr); i++ {
		arr[i] = float32(i)
	}
	//fmt.Print(arr)
	fmt.Println(Sum(arr))
	fmt.Println(SumAndAverage(2, 3))
}
func Sum(arrF []float32) float32 {
	var sum float32 = 0.0
	for i := 0; i < len(arrF); i++ {
		sum += arrF[i]
	}
	return sum
}

func SumAndAverage(n1 int, n2 int) (int, float32) {
	return n1 + n2, float32((n1 + n2) / 2)
}
