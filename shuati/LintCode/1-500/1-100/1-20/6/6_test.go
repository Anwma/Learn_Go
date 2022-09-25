package main

import (
	"sort"
	"testing"
)

func MergeSortedArray(a []int, b []int) []int {
	// write your code here
	a = append(a, b...)
	sort.Ints(a)
	return a
}

func MergeSortedArray1(a []int, b []int) []int {
	// write your code here
	//1.双指针
	res := make([]int, len(a)+len(b))
	ptrA, ptrB := 0, 0
	index := 0
	for ptrA < len(a) && ptrB < len(b) {
		if a[ptrA] < b[ptrB] {
			res[index] = a[ptrA]
			ptrA++
			index++
		} else {
			res[index] = b[ptrB]
			ptrB++
			index++
		}
	}
	for ptrA < len(a) {
		res[index] = a[ptrA]
		ptrA++
		index++
	}
	for ptrB < len(b) {
		res[index] = b[ptrB]
		ptrB++
		index++
	}
	return res
}

func BenchmarkMergeSortedArray(b *testing.B) {
	arrA := []int{1, 2, 3, 4}
	arrB := []int{2, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		MergeSortedArray(arrA, arrB)
		//0		121.4 ns/op
		//1		39.71 ns/op
	}
}
