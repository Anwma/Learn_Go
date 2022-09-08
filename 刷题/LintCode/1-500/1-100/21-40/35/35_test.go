package main

import "testing"

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
		//cur.Next,pre,cur = pre,cur,cur.Next
	}
	return pre
}
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//Reverse()
	}
}
