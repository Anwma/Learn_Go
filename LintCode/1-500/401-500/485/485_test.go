package _85

import "testing"

func Generate(size int) []int {
	// write your code here
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i + 1
	}
	return slice

	//arry := make([]int, 0, 0)
	//for i := 1; i <= size; i++ {
	//	arry = append(arry, i)
	//}
	//return arry
}
func BenchmarkGenerate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Generate(10)
	}
	b.StopTimer()
}
