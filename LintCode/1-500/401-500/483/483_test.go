package _83

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

func ToArrayList(head *ListNode) []int {
	// write your code here
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}



func BenchmarkToArrayList(b *testing.B) {
	list := []int{3, 5, 8}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToArrayList(head)
	}
	b.StopTimer()
}
