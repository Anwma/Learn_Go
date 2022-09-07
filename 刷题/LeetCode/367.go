package main

import "math"

func isPerfectSquare(num int) bool {
	x := int(math.Sqrt(float64(num)))
	return x*x == num
}

func isPerfectSquare2(num int) bool {
	for x := 1; x*x <= num; x++ {
		if x*x == num {
			return true
		}
	}
	return false
}

func isPerfectSquare3(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square < num {
			left = mid + 1
		} else if square > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func isPerfectSquare4(num int) bool {
	x0 := float64(num)
	for {
		x1 := (x0 + float64(num)/x0) / 2
		if x0-x1 < 1e-6 {
			x := int(x0)
			return x*x == num
		}
		x0 = x1
	}
}


