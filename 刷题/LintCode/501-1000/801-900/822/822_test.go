package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseStore(head *ListNode) []int {
	// write your code here
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	//反转切片
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (h *ListNode) Append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

func BenchmarkReverseStore(b *testing.B) {
	list := []int{2,4,3}
	head := &ListNode{Val: list[0]}
	for i:=1;i<len(list);i++ {
		head.Append(list[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseStore(head)
	}
	b.StopTimer()
}
