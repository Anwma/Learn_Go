package _37

import "math"

func CopyBooks(pages []int, k int) int {
	// write your code here
	if pages == nil || len(pages) == 0 {
		return 0
	}
	if k == 0 {
		return -1
	}
	start, end := 0, getSum(pages)

	for start+1 < end {
		mid := start + (end-start)>>1
		if getNumberOfCopiers(pages, mid) <= k {
			end = mid
		} else {
			start = mid
		}
	}
	if getNumberOfCopiers(pages, start) <= k {
		return start
	}
	return end
}

func getNumberOfCopiers(pages []int, limit int) int {
	copiers, lastCopied := 0, limit

	for _, page := range pages {
		if page > limit {
			return math.MaxInt32
		}
		if lastCopied+page > limit {
			copiers++
			lastCopied = 0
		}
		lastCopied += page
	}
	return copiers
}

func getSum(pages []int) int {
	sum := 0
	for _, page := range pages {
		sum += page
	}
	return sum
}
