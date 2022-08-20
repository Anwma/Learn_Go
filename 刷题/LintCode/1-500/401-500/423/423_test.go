package main

func IsValidParentheses(s string) bool {
	// write your code here
	//1.corner case 排除s元素个数是单数的情况
	n := len(s)
	if n%2 != 0 {
		return false
	}
	//2.建立一个hashmap映射
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//3.
	res := []byte{}
	for _, v := range s {
		if m[byte(v)] > 0 {
			if len(res) == 0 || m[byte(v)] != res[len(res)-1] {
				res = append(res, byte(v))
			} else {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, byte(v))
		}
	}
	return len(res) == 0
}
