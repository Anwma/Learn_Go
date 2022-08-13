package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// ! 不会被打印出来
	defer fmt.Println("!")

	os.Exit(3)
}

func twoSum(nums []int, target int) []int {
	//1.暴力解法
	// for i:=0;i<len(nums);i++{
	//     for j := 0;j<len(nums);j++{
	//         if nums[i]+nums[j] == target && i!= j{
	//             return []int{i,j}
	//         }
	//     }
	// }
	// return nil
	//2.哈希表
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		other := target - nums[i]
		if p, ok := m[other]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}
func isPalindrome(s string) bool {
	q, p := 0, len(s)-1
	s = strings.ToUpper(s)
	for q <= p {
		//确保左右两边都是有效字符
		if !ischar(s[q]) {
			q++
		} else if !ischar(s[p]) {
			p--
		} else if s[q] != s[p] {
			return false
		}
		//是字符，正常左右两边走
		q++
		p--
	}
	return true

}
func ischar(ch byte) bool {
	return (ch <= 'Z' && ch >= 'A') || (ch >= '0' && ch <= '9')
}
