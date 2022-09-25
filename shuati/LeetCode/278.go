package main

func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m - 1
			continue
		}
		l = m + 1
	}
	return l
}

func isBadVersion(flag int) bool {
	return true
}
