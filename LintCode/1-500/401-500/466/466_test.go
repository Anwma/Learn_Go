package _66

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

func CountNodes(head *ListNode) int {
	// write your code here
	count := 0
	for head != nil {
		count ++
		head = head.Next
	}
	return count
}

func BenchmarkCountNodes(b *testing.B) {
	list := []int{3, 5, 8}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountNodes(head)
	}
	b.StopTimer()
}
