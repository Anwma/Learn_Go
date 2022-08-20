package main

type ListNode struct {
	Val int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	// write your code here
	src,slow,fast := head,head,head.Next
	for slow != nil || fast != nil{
		slow.Val,fast.Val = fast.Val,slow.Val
		slow = fast.Next
		fast = slow.Next
	}
	return src
}
