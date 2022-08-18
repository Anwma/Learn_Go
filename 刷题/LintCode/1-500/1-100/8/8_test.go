package main

func RotateString(s []byte, offset int) {
	// write your code here
	if offset == 0 || len(s) == 0 {
		return
	}
	reverse(s, 0, offset)
	reverse(s, offset+1, len(s)-1)
	reverse(s, 0, len(s)-1)
}
func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
