package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
func AppendByte(slice []byte, data ...byte) []byte {
	//获取目标数组的len
	m := len(slice)
	//获取total-len
	n := m + len(data)
	//若 total-len > cap dst 数组
	if n > cap(slice) {
		// if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		//将原slice copy到 det slice 钟
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
