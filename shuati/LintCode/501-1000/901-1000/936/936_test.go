package _36

func CapitalizesFirst(s string) string {
	// Write your code here
	n := len(s)
	sStr := []byte(s)
	if sStr[0] >= 'a' && sStr[0] <= 'z' {
		sStr[0] -= 32
	}
	for i := 0; i < n; i++ {
		//当前字符的前面一个字符为' '且当前字符不为' '
		if sStr[i-1] == ' ' && sStr[i] != ' ' {
			sStr[i] -= 32
		}
	}
	return string(sStr)
}
