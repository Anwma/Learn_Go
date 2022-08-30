package main

func MinWindow(source string, target string) string {
	// write your code here
	// corner case
	if len(source) == 0 {
		return ""
	}
	//窗口中出现的字符次数
	countS := make([]int, 128)
	countT := make([]int, 128)
	//记录不同字符的个数
	k := 0
	for i := 0; i < len(target); i++ {
		countT[target[i]]++
		if countT[target[i]] == 1 {
			k++
		}
	}
	c := 0
	ansLeft, ansRight := -1, -1
	left, right := 0, 0
	//left : 主指针
	//right: 辅助指针
	for left < len(source) {
		for right < len(source) && c < k {
			countS[source[right]]++
			if countS[source[right]] == countT[source[right]] {
				c++
			}
			right++
		}
		if c == k {
			if ansLeft == -1 || right-left < ansRight-ansLeft {
				ansLeft = left
				ansRight = right
			}
		}
		countS[source[left]]--
		if countS[source[left]] == countT[source[left]]-1 {
			c--
		}
		left++
	}
	if ansLeft == -1 {
		return  ""
	}else {
		return source[ansLeft:ansRight]
	}
}
