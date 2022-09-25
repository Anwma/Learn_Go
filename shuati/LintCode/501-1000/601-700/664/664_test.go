package main

func main() {

}

func CountBits(num int) []int {
	// write your code here
	f := make([]int, num+1)
	for i := 1; i <= num; i++ {
		f[i] = f[i>>1] + (i % 2)
	}
	return f
}
