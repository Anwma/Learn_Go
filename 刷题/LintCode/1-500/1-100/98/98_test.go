package _8

import (
	"sort"
	"testing"
)

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

func SortList(head *ListNode) *ListNode {
	// write your code here
	if head == nil || head.Next == nil{
		return head
	}
	arr := []int{}
	//1.先转换为数组
	tmp := head
	for tmp != nil {
		arr = append(arr, tmp.Val)
		tmp = tmp.Next
	}
	//2.对数组排序
	sort.Ints(arr)
	//3.把数组放回到链表当中
	newArr := &ListNode{Val: arr[0]}
	for i := 1; i < len(arr); i++ {
		newArr.Append(arr[i])
	}
	return newArr
}

func BenchmarkSortList(b *testing.B) {
	list := []int{4, 5, 1, 2, 3}
	head := &ListNode{Val: list[0]}
	for i := 1; i < len(list); i++ {
		head.Append(list[i])
	}

	for i := 0; i < b.N; i++ {
		SortList(head)
	}
}
