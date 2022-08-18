package main

func PrimeFactorization(num int) []int {
	// write your code here
	res := []int{}
	//num=8      i<= 3
	for i := 2; i <= num/i; {
		if num%i == 0 {
			res = append(res, i) //[2]
			num = num / i        //num=[4]
		} else {
			i++
		}
	}
	if num != 1 {
		res = append(res, num)
	}
	return res
}

//func isPrime(n int) bool {
//	if n == 0 || n == 1 {
//		return false
//	}
//	for i := 2; i <= n/i; i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}
func main() {
	PrimeFactorization(7)
}
