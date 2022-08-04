package main

import "fmt"

func main() {
	//print(CharToInteger('a'))
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//newArr := Rotate(nums, 3)
	//fmt.Println(newArr)
	//res := Prime(1)
	//fmt.Println(res)
	//start := time.Now()
	//for i := 0; i < 10000000; i++ {
	//	IsPalindrome(4294967296 - 1)
	//}
	//end := time.Since(start)
	//fmt.Println(end)
	// write your code here
	//res := CalcYangHuisTriangle(4)
	//fmt.Println(res)
	nums := []int{1, 0, 3, 5, 0}
	MoveZeroes(nums)
	fmt.Println(nums)

}

func MoveZeroes(nums []int) {
	// write your code here
	start := 0
	for i := 0; i < len(nums); i++ {
		//碰到当前的值为0 则将其移动到最前面
		if nums[i] == 0 {
			nums[i], nums[start] = nums[start], nums[i]
			start++
		}
		//fmt.Println(nums)
	}
	nums = append(nums[start:], nums[:start]...)
}

//func CalcYangHuisTriangle(n int) [][]int {
//	//write your code here
//	tmp := []int{1}
//	res := [][]int{}
//	for i := 0; i < n; i++ {
//		res = append(res, tmp)
//		for j := 1; j < i; j++ {
//			res[i] = append(res[i], res[i-1][j-1]+res[i-1][j])
//		}
//		if i > 0 {
//			res[i] = append(res[i], 1)
//		}
//	}
//	return res
//}

//func IsPalindrome(n int) bool {
//	// Write your code here
//	if n <= 1 {
//		return true
//	}
//	a := toBinary(n)
//	for i := 0; i <= len(a)-i-1; i++ {
//		if a[i] == a[len(a)-i-1] {
//			continue
//		} else {
//			return false
//		}
//	}
//	return true
//}
//
//func toBinary(n int) []int {
//	a := make([]int, 0,33)
//	for {
//		a = append(a, n%2)
//		n = n / 2
//		if n == 1 {
//			a = append(a, 1)
//			break
//		}
//	}
//	return a
//}

//func Prime(n int) []int {
//	// write your code here
//	//slice := []int{}
//	var slice []int
//	fmt.Println("初始化的",slice)
//	for i := 2; i <= n; i++ {
//		if isPrime(i) {
//			slice = append(slice, i)
//		}
//	}
//	return slice
//}
//func isPrime(n int) bool {
//	if n < 2 {
//		return false
//	}
//	for i := 2; i <= n/i; i++ {
//		if n%i == 0 {
//			return false
//		}
//	}
//	return true
//}

//func PrintX(n int) []string {
//	// write your code here
//	//1.创建一个长度为 n 的 []string
//	res := make([]string, n)
//	//2.针对每一个[]里面,对其进行操作
//	for i := 0; i < n; i++ {
//		//3.每一个切片中,开辟大小为n的byte[]
//		item := make([]byte, n)
//		//4.每个[]byte中的item初始化为 ' '
//		for j := range item {
//			item[j] = ' '
//		}
//		//5.只需对当前的行切片[]byte的首尾两个index赋值为 'X' 即可
//		item[i] = 'X'
//		item[n-i-1] = 'X'
//		//6.对 []string 赋相应的 []byte
//		res[i] = string(item)
//	}
//	//7.返回结果
//	return res
//}

//func SecondMax(nums []int) int {
//	// write your code here
//	sort.Ints(nums)
//	return nums[len(nums)-2]
//}

//func ReverseArray(nums []int) {
//	// write your code here
//	n := len(nums) - 1
//	for i := 0; i < len(nums)/2; i++ {
//		nums[i], nums[n-i] = nums[n-i], nums[i]
//	}
//}

//func MaxNum(nums []int) int {
//	// write your code here
//	sort.Ints(nums)
//	return nums[len(nums)-1]
//}

//func MaxOfArray(a []float32) float32 {
//	// write your code here
//	max := a[0]
//	for i := 1; i < len(a); i++ {
//		if a[i] > max {
//			max = a[i]
//		}
//	}
//	return max
//}

//func SortIntegers(a []int)  {
//	// write your code here
//	sort.Ints(a)
//}

//func SwapIntegers(a []int, index1 int, index2 int) {
//	// write your code here
//	a[index1], a[index2] = a[index2], a[index1]
//}

//func IsAlphanumeric(c byte) bool {
//	// write your code here
//	if c <= 'Z' && c >= 'A' || c <= 'z' && c >= 'a' || c <= '9' && c >= '0' {
//		return true
//	}
//	return false
//}

//func GetTheMonthDays(year int, month int) int {
//	// write your code here
//	if month == 2 {
//		if IsLeapYear(year) {
//			return 29
//		}
//		return 28
//	}
//	if month == 4 || month == 6 || month == 9 || month == 11 {
//		return 30
//	}
//	return 31
//}
//
//func IsLeapYear(n int) bool {
//	// write your code here
//	if n%4 == 0 && n%100 != 0 || n%400 == 0 {
//		return true
//	}
//	return false
//}

//func LowercaseToUppercase(character byte) byte {
//	// write your code here
//	return character-32
//}

//func MaxOfThreeNumbers(num1 int, num2 int, num3 int) int {
//	// write your code here
//	tmp := max(num1, num2)
//	return max(tmp, num3)
//}
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func Calculate(a int, op byte, b int) int {
//	// write your code here
//	res := 0
//	switch op {
//	case '+':
//		res = a + b
//	case '-':
//		res = a - b
//	case '*':
//		res = a * b
//	case '/':
//		res = a / b
//	}
//	return res
//}

//func Calculate(r int) []float64 {
//	// write your code here
//	//PI := 3.14
//	//return []float64{float64(r) * 2 * PI, float64(r*r) * PI}
//}

//func ReverseInteger(number int) int {
//	// write your code here
//	return number%10*100 + number/10%10*10 + number/100
//}

//func Rotate(nums []int, k int) []int {
//	//根据k切分成两部分,k可能大于nums的长度,故取模
//	n := len(nums) - 1
//	return append(nums[n-k%n:], nums[:n-k%n]...)
//}

//func CanPlaceFlowers(flowerbed []int, n int) bool {
//	// Write your code here
//	length := len(flowerbed)
//	for i, v := range flowerbed {
//		//首尾部 当前值为0,且首部下一个index==0,尾部上一个index==0。可进行放置
//		//中部 当前值为0 其前面一个和后面一个index均为0 方可放置
//		if v == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
//			n--
//			//并将当前值置为1
//			flowerbed[i] = 1
//			if n <= 0 {
//				return true
//			}
//		}
//	}
//	//n 为 0的情况
//	return n == 0
//}

//func SumofTwoStrings(a string, b string) string {
//	// write your code here
//	res := ""
//	indexA, indexB := len(a)-1, len(b)-1
//	for indexA >= 0 && indexB >= 0 {
//		temp := a[indexA] - '0' + b[indexB] - '0'
//		res = strconv.Itoa(int(temp)) + res
//		indexA--
//		indexB--
//	}
//	if indexA >= 0 {
//		res = a[0:indexA] + res
//	}
//	if indexB >= 0 {
//		res = a[0:indexB] + res
//	}
//	return res
//}

//func MinimumSize(nums []int, s int) int {
//	// write your code here
//	n := len(nums)
//	answer := math.MaxInt64
//	for start := 0; start < n; start++ {
//		for end := 0; end < n; end++ {
//			sum := 0
//			for i := start; i <= end; i++ {
//				sum += nums[i]
//			}
//			if sum >= s{
//				answer = int(math.Min(float64(answer), float64(end-start+1)))
//			}
//		}
//	}
//	return answer
//}

/**
 * @param str: the input string
 * @return: The lower case string
 */

//func ToLowerCase(str string) string {
//	// Write your code here
//	return strings.ToLower(str)
//}

/**
 * @param size: An integer
 * @return: An integer list
 */

//func Generate(size int) []int {
//	// write your code here
//	slice := make([]int, size)
//	for i := 0; i < size; i++ {
//		slice[i] = i + 1
//	}
//	return slice
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * @param head: the head of linked list.
 * @return: An integer list
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func ToArrayList(head *ListNode) []int {
//	// write your code here
//	slice := make([]int,0)
//	for head != nil{
//		slice = append(slice,head.Val)
//		head = head.Next
//	}
//	return slice
//}

/**
 * @param character: a character
 * @return: An integer
 */

//func CharToInteger(character byte) int {
//	// write your code here
//	return int(character)
//}
