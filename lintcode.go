package main

import (
	"strconv"
)

func main() {
	//print(CharToInteger('a'))
}
func SumofTwoStrings(a string, b string) string {
	// write your code here
	res := ""
	indexA, indexB := len(a)-1, len(b)-1
	for indexA >= 0 && indexB >= 0 {
		temp := a[indexA] - '0' + b[indexB] - '0'
		res = strconv.Itoa(int(temp)) + res
		indexA--
		indexB--
	}
	if indexA >= 0 {
		res = a[0:indexA] + res
	}
	if indexB >= 0 {
		res = a[0:indexB] + res
	}
	return res
}

//func MinimumSize(nums []int, s int) int {
//	// write your code here
//	n := len(nums)
//	answer := math.MaxInt64
//	for start := 0; start < n; start++ {
//		for end := 0; end < n; end++ {
//			sum := 0
//			for i := start; i <= end; i++ {
//				sum += nums[i]
//			}
//			if sum >= s{
//				answer = int(math.Min(float64(answer), float64(end-start+1)))
//			}
//		}
//	}
//	return answer
//}

/**
 * @param str: the input string
 * @return: The lower case string
 */

//func ToLowerCase(str string) string {
//	// Write your code here
//	return strings.ToLower(str)
//}

/**
 * @param size: An integer
 * @return: An integer list
 */

//func Generate(size int) []int {
//	// write your code here
//	slice := make([]int, size)
//	for i := 0; i < size; i++ {
//		slice[i] = i + 1
//	}
//	return slice
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * @param head: the head of linked list.
 * @return: An integer list
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func ToArrayList(head *ListNode) []int {
//	// write your code here
//	slice := make([]int,0)
//	for head != nil{
//		slice = append(slice,head.Val)
//		head = head.Next
//	}
//	return slice
//}

/**
 * @param character: a character
 * @return: An integer
 */

//func CharToInteger(character byte) int {
//	// write your code here
//	return int(character)
//}
