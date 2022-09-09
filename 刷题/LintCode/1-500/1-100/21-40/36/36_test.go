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

func ReverseBetween(head *ListNode, left, right int) *ListNode {
	// 建立虚拟节点 当left=0时 虚拟节点能避免做过多处理
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next     //next = 5
		cur.Next = next.Next //cur.Next = 4
		next.Next = pre.Next //next.Next = 2
		pre.Next = next      //pre.Next = 5
		//cur.Next, cur.Next.Next, pre.Next = cur.Next.Next, pre.Next, cur.Next
	}
	return dummyNode.Next
}

func BenchmarkReverse(b *testing.B) {
	list := []int{9, 7, 2, 5, 4, 3, 6}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}
	for i := 0; i < b.N; i++ {
		ReverseBetween(head, 2, 5)
	}
}
