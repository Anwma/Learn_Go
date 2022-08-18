package _19
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

func InsertNode(head *ListNode, val int) *ListNode {
	// write your code here


	return nil
}