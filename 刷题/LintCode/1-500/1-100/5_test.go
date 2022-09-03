package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func KthLargestElement(k int, nums []int) int {
	// write your code here
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func KthLargestElement1(k int, nums []int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func KthLargestElement2(k int, nums []int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums) - k + 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize/2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i * 2 + 1, i * 2 + 2, i
	if l < heapSize && a[l] > a[largest] {
		largest = l
	}
	if r < heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}
func KthLargestElement3(k int, nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	res := KSelect(nums,0,len(nums)-1,k-1)
	return res
}

func KSelect(nums []int,start int,end int,k int) int {
	if start == end {
		return nums[start]
	}

	left,right := start,end
	povit := nums[start+ (end - start)/2]
	for left <= right {
		for left <= right && nums[left] > povit{
			left++
		}
		for left <= right && nums[right] < povit{
			right--
		}
		if left <= right {
			nums[left],nums[right]=nums[right],nums[left]
			left++
			right--
		}

	}
	if k >= start && k <= right {
		return KSelect(nums,start,right,k)
	}
	if k >= left && k <= end {
		return  KSelect(nums,left,end,k)
	}
	return nums[right+1]
}
func BenchmarkKthLargestElement(b *testing.B) {
	nums := []int{1, 5, 4, 8, 2}
	for i := 0; i < b.N*10; i++ {
		KthLargestElement2(1, nums)
		//0  68.94 ns/op
		//1  8705 ns/op
		//2  18.34 ns/op
		//3  19.34 ns/op
	}
}
