package _25

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

func FindNode(head *ListNode, val int) *ListNode {
	// write your code here
	for head != nil {
		if head.Val == val{
			return head
		}
		head = head.Next
	}
	return nil
}

func BenchmarkFindNode(b *testing.B) {
	list := []int{1,2,3,4,5,6,7,8,9,10}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindNode(head,2)
	}
	b.StopTimer()
}