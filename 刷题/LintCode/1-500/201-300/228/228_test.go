package _28

import (
	"testing"
)

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
	//BenchmarkMiddleNode-8            2439154               483.7 ns/op
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	count := 0
	m := map[int]*ListNode{}
	for head != nil {
		count++
		m[count] = head
		head = head.Next
	}
	if count%2 == 0 {
		return m[count/2]
	} else {
		return m[count/2+1]
	}



	//BenchmarkMiddleNode-8           169528088                6.928 ns/op
	//if head == nil {
	//	return head
	//}
	//node := head
	//count := 0
	//for head.Next != nil {
	//	count++
	//	head = head.Next
	//	if count%2 == 0 {
	//		node = node.Next
	//	}
	//}
	//return node

	//ME
	//BenchmarkMiddleNode-8           397173186                2.987 ns/op
	//if head == nil || head.Next == nil || head.Next.Next == nil {
	//	return head
	//}
	//pre := head
	//cur := head.Next.Next
	//for cur.Next != nil && cur.Next.Next != nil{
	//	pre = pre.Next
	//	cur = cur.Next.Next
	//}
	//return pre.Next
}

func BenchmarkMiddleNode(b *testing.B) {
	list := []int{1,2,3,4,5,6,7,8,9,10}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MiddleNode(head)
	}
	b.StopTimer()
	//BenchmarkMiddleNode-8           28508912                44.23 ns/op
}
