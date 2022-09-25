package main

import "fmt"

func LongestLightingTime(operation [][]int) byte {
	// write your code here
	char := []byte{
		'a', 'b', 'c', 'd', 'e', 'f',
		'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'q', 'r',
		's', 't', 'u', 'v', 'w', 'x',
		'y', 'z',
	}
	//记录每一盏灯亮起的时间
	time := make([]int, len(char))
	temp := make([]int, len(char))
	//第n栈灯的时间为
	time[operation[0][0]] = operation[0][1]
	for i := 1; i < len(operation); i++ {
		temp[operation[i][0]] = operation[i][1] - operation[i-1][1]
		if temp[operation[i][0]] > time[operation[i][0]] {
			time[operation[i][0]] = temp[operation[i][0]]
		}
	}
	//查找时间最长的值 单次亮起的时间最长
	index := 0
	max := time[0]
	for i := 1; i < len(time); i++ {
		if time[i] > max {
			max = time[i]
			index = i
		}
	}
	return char[index]
}
func main() {
	operation := [][]int{{0, 2}, {1, 5}, {0, 9}, {2, 15}}
	fmt.Printf("%c", LongestLightingTime(operation))

}
