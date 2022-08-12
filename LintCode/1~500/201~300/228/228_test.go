package _28
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