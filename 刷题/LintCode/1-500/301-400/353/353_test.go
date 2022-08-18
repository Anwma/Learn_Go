package _53

func LargestLetter(s string) string {
	// write your code here
	n := len(s)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		k := key(s[i])
		m[int(k)] = 1
	}
	return ""
}
func key(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 'a'
	}else {
		return 26 + c - 'A'
	}
}
