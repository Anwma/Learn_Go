package _60

/**
 * @param a: an integer array
 * @param target: An integer
 * @param k: An integer
 * @return: an integer array
 */
func KClosestNumbers(a []int, target int, k int) []int {
	// write your code here
	left := findLowerClosest(a, target)
	right := left + 1

	result := make([]int, k)
	for i := 0; i < k; i++ {
		if isLeftCloser(a, target, left, right) {
			result[i] = a[left]
			left--
		} else {
			result[i] = a[right]
			right++
		}
	}
	return result
}

func findLowerClosest(a []int, target int) int {
	l, r := 0, len(a)-1
	for l+1 < r {
		m := l + (r-l)>>1
		if a[m] < target {
			l = m
		} else {
			r = m
		}
	}

	if a[r] < target {
		return r
	}
	if a[l] < target {
		return l
	}
	return -1
}

func isLeftCloser(a []int, target, left, right int) bool {
	if left < 0 {
		return false
	}
	if right >= len(a) {
		return true
	}
	if target-a[left] != a[right]-target {
		return target-a[left] < a[right]-target
	}
	return true
}
