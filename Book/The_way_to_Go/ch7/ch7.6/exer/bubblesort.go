// 7.16
package main

import "fmt"

func main() {
	sliceInt := []int{2, 1, 4, 5, 8, 6}
	fmt.Print(bubbleSort(sliceInt))
}
func bubbleSort(tar []int) []int {
	lens := len(tar)
	//冒泡排序 从小到大
	for i := 0; i < lens; i++ {
		for j := i; j < lens; j++ {
			if tar[i] > tar[j] {
				tar[i], tar[j] = tar[j], tar[i]
			}
		}
	}
	return tar
}
