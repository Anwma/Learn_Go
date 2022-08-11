package _343

import (
	"strconv"
	"testing"
)

func SumofTwoStrings(a string, b string) string {
	// write your code here
	res := ""
	indexA, indexB := len(a)-1, len(b)-1
	for indexA >= 0 && indexB >= 0 {
		temp := a[indexA] - '0' + b[indexB] - '0'
		res = strconv.Itoa(int(temp)) + res
		indexA--
		indexB--
	}
	if indexA >= 0 {
		res = a[0:indexA+1] + res
	}
	if indexB >= 0 {
		res = b[0:indexB+1] + res
	}
	return res
}
func BenchmarkSumofTwoStrings(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumofTwoStrings("99", "111")
	}
	b.StopTimer()

}
