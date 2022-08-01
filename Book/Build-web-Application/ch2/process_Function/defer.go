package main

import "fmt"

/*

	func ReadWrite() bool {
		file.Open("file")
		// 做一些工作
		if failureX {
			file.Close()
			return false
		}

		if failureY {
			file.Close()
			return false
		}

		file.Close()
		return true
	}

*/

/*

	func ReadWrite() bool {
		file.Open("file")
		defer file.Close()
		if failureX {
			return false
		}
		if failureY {
			return false
		}
		return true
	}

*/

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
