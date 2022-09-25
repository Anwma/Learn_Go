package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) Append(i int) {
	//往链表中添加元素
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

func MiddleNode(head *ListNode) *ListNode {
	// write your code here.
	//corner case
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//count := 0
	//res := head
	//for head != nil {
	//	count++
	//	if count%2 == 0 {
	//		res = res.Next
	//	}
	//	head = head.Next
	//}
	//return res

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func BenchmarkMiddleNode(b *testing.B) {
	list := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}

	for i := 0; i < b.N; i++ {
		MiddleNode(head)
	}
}
