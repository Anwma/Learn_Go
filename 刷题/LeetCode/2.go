package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := ListNode{
		Val:  2,
		Next: nil,
	}
	fmt.Printf("%v", addTwoNumbers(&l1, &l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.创建一个新的链表
	res := &ListNode{}
	ans := res
	//2.对链表中的值进行相加
	overTen := false
	for l1 != nil || l2 != nil {
		//创建一个临时值 对二者的值相加
		if l1 != nil {
			res.Val += l1.Val
		}
		if l2 != nil {
			res.Val += l2.Val
		}
		if overTen {
			res.Val++
			overTen = false
		}
		if res.Val >= 10 {
			res.Val = res.Val % 10
			overTen = true
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil || l2 != nil {
			res.Next = &ListNode{}
			res = res.Next
		}
	}
	if overTen {
		res.Next = &ListNode{Val: 1}
	}
	return ans
}

//&{4 <nil>}  不理解 ？？？ 为啥跑不通...

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//1.创建一个新的链表
	res := &ListNode{}
	ans := res
	//2.对链表中的值进行相加
	add := 0
	for l1 != nil && l2 != nil {
		//创建一个临时值 对二者的值相加
		res.Val = l1.Val + l2.Val
		//如果值 > 10 产生进位 最多为一
		if res.Val/10 == 1 {
			add = 1
		}
		res.Val = res.Val % 10
		//链表往后移动一位
		res = res.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		res.Val = l1.Val%10 + add
		if (l1.Val+add)/10 == 1 {
			add = 1
		}
		res = res.Next
		l1 = l1.Next
	}
	for l2 != nil {
		res.Val = l2.Val%10 + add
		if (l2.Val+add)/10 == 1 {
			add = 1
		}
		res = res.Next
		l2 = l2.Next
	}
	if add == 1 {
		ans.Next = &ListNode{Val: add}
	}
	return ans
}

 */
