package main

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

func Reverse(head *ListNode) *ListNode {
	// write your code here
	var pre *ListNode
	cur := head
	for cur != nil {
		//temp := cur.Next
		//cur.Next = pre
		//pre = cur
		//cur = temp
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
func BenchmarkReverse(b *testing.B) {
	list := []int{2, 4, 3, 1}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}
	for i := 0; i < b.N; i++ {
		Reverse(head)
	}
}
