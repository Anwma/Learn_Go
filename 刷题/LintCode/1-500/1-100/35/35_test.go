package _5

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {
	// write your code here
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
