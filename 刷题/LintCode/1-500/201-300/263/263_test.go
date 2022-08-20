package main

func MatchParentheses(string string) bool {
	// write your code here
	count := 0
	for _, s := range string {
		if s == '(' {
			count++
		}
		if s == ')' {
			count--
		}
		//若先进来一个 ')' 导致count < 0 可直接返回false
		if count < 0 {
			return false
		}
	}
	// ( ) 正负相消的时候 也就是count == 0 括号正好匹配
	if count == 0 {
		return true
	}
	return false
}
