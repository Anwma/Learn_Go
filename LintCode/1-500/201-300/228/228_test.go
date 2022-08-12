package _28

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) Append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

func MiddleNode(head *ListNode) *ListNode {
	// write your code here


	return nil
}

func BenchmarkMiddleNode(b *testing.B) {

	for i := 0; i < b.N; i++ {

	}
}