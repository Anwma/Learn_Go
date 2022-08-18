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

func HasCycle(head *ListNode) bool {
	// write your code here
	//nil
	if head == nil || head.Next == nil{
		return false
	}
	slow, fast := head, head.Next
	//因为要判断这两个指针是否相遇,所以fast为 head.Next
	for slow != fast{
		//1 2
		if fast == nil || fast.Next == nil{
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
func BenchmarkHasCycle(b *testing.B) {
	list := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}
	for i := 0; i < b.N; i++ {
		HasCycle(head)
	}
}
