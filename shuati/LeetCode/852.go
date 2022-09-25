package main

func peakIndexInMountainArray(arr []int) int {
	//只需找到临界点 该点前后的二值均小于该值
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
			return i
		}
	}
	return -1
}

func peakIndexInMountainArray2(a []int) int {
	// Write your code here
	l, r := 1, len(a)-1
	for l+1 < r {
		m := l + (r-l)>>1
		if a[m] > a[m-1] && a[m] > a[m+1] {
			return m
		} else if a[m] < a[m-1] && a[m] > a[m+1] {
			//右侧
			r = m
		} else {
			l = m
		}
	}
	if a[l] > a[l-1] && a[l] > a[l+1] {
		return l
	}
	if a[r] > a[r-1] && a[r] > a[r+1] {
		return r
	}
	return -1
}
