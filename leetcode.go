package main

import "fmt"

/**
 * @param nums: The integer array.
 * @param target: Target to find.
 * @return: The first position of target. Position starts from 0.
 */
//func BinarySearch(nums []int, target int) int {
//	// write your code here
//	l,r := 0,len(nums)-1
//	m := l + (r - l) >> 1
//	for l + 1 < r{
//		//这边必须是target > nums[m]
//		if target > nums[m]{
//			l = m
//		}else {
//			r = m
//		}
//		m = l + (r - l) >> 1
//	}
//	if nums[l] == target {
//		return l
//	}
//	if nums[r] == target {
//		return r
//	}
//	return -1
//}

/**
 * @param s: input string
 * @return: a string as the longest palindromic substring
 */

//func LongestPalindrome(s string) string {
//	// write your code here
//	for length := len(s); length > 0; length-- {
//		for start := 0; start+length <= len(s); start++ {
//			if isPalindrome(s, start, start+length-1) {
//				return s[start : start+length]
//			}
//		}
//	}
//	return ""
//}
//func isPalindrome(s string, left, right int) bool {
//	for left < right && s[left] == s[right] {
//		left++
//		right--
//	}
//	return left >= right
//}

// LongestPalindrome 解法一：
//func LongestPalindrome(s string) string {
//	// write your code here
//	for lens := len(s); lens >= 1; lens-- {
//		for i := 0; i+lens <= len(s); i++ {
//			l, r := i, i+lens-1
//			for l < r && s[l] == s[r] {
//				l++
//				r--
//			}
//			if l >= r {
//				return s[i:i+lens]
//			}
//		}
//	}
//	return ""
//}
//func LongestPalindrome(s string) string {
//	// write your code here
//	longest, start := 0, 0
//	for i := 0; i < len(s); i++ {
//		var left, right int
//		//odd 奇数情况
//		left, right = i, i
//		for left >= 0 && right < len(s) && s[left] == s[right] {
//			left--
//			right++
//		}
//		if longest < right-left-1 {
//			longest = right - left - 1
//			start = left + 1
//		}
//		//偶数情况
//		left, right = i, i+1
//		for left >= 0 && right < len(s) && s[left] == s[right] {
//			left--
//			right++
//		}
//		if longest < right-left-1 {
//			longest = right - left - 1
//			start = left + 1
//		}
//	}
//	return s[start : start+longest]
//}
//解法四：中心线枚举算法 Coding Quality
func LongestPalindrome(s string) string {
	// write your code here
	longest := ""
	for i := 0; i < len(s); i++ {
		oddPalindrome := getPalindromeFrom(s, i, i)
		if len(longest) < len(oddPalindrome) {
			longest = oddPalindrome
		}
		evenPalindrome := getPalindromeFrom(s, i, i+1)
		if len(longest) < len(evenPalindrome) {
			longest = evenPalindrome
		}
	}
	fmt.Println("这里是return")
	return longest
}
func getPalindromeFrom(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
func main() {
	//nums := []int{3, 4, 5, 8, 8, 8, 8, 10, 13, 14}
	//fmt.Println(BinarySearch(nums, 8))
	//a := AreaOfTrapezoid(12,3,5)
	//fmt.Println(a)
	//fmt.Println(isPrime(5))
	//isPrime(5)
	//Prime(5)
	//fmt.Println(Prime(5))
	//carrot := [][]int{
	//	{5, 3, 7, 1, 7},
	//	{4, 6, 5, 2, 8},
	//	{2, 1, 1, 4, 6},
	//}
	//carrot := [][]int{
	//	{5, 7, 6, 3},
	//	{2, 4, 8, 12},
	//	{3, 5, 10, 7},
	//	{4, 16, 4, 17},
	//}
	//math.pow
	//fmt.Println(PickCarrots(carrot))
	fmt.Println(LongestPalindrome("abb"))
}

/**
 * @param n: The number of digits
 * @return: All narcissistic numbers with n digits
 */
//import "fmt"
//
//func GetNarcissisticNumbers(n int) []int {
//	// write your code here
//	res := make([]int, 0)
//	//first 表示n位数的开始值
//	first := pow(10, n-1)
//	//last 表示n位数的终止值
//	last := first * 10
//	//添加0
//	if n == 1 {
//		res = append(res, 0)
//	}
//	for i := first; i < last; i++ {
//		if isNarcissisticNumbers(i) {
//			res = append(res, i)
//		}
//	}
//	return res
//}
//
////自构pow函数,而非使用math.pow
//func pow(base, pows int) int {
//	res := 1
//	for i := 0; i < pows; i++ {
//		res *= base
//	}
//	return res
//}
//
////判断某个数是否是水仙花数
//func isNarcissisticNumbers(n int) bool {
//	sum := 0
//	tmp := n
//	//利用Sprint(return - string)得到输入n的长度,也就是幂次
//	pows := len(fmt.Sprint(n))
//	for tmp >= 1 {
//		sum += pow(tmp%10, pows)
//		tmp /= 10
//	}
//	return sum == n
//}

//
//func Sqrt(x int) int {
//	// write your code here
//	if x <= 1 {
//		return x
//	}
//	var c, x0 float64 = float64(x), float64(x)
//	for {
//		var xi = 0.5 * (x0 + c/x0)
//		if math.Abs(x0-xi) < 1e-7 {
//			break
//		}
//		x0 = xi
//	}
//	return int(x0)
//}

//func PickCarrots(carrot [][]int) int {
//	// write your code here
//	//1.从矩阵的中心点出发 (m,n)
//	m, n := (len(carrot)-1)/2, (len(carrot[0])-1)/2
//	//2.标记某些位置是否走过,走过之后将其置为0
//	carrotNum := carrot[m][n]
//	carrot[m][n] = 0
//	//定义 上 下 左 右
//	road := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
//	flag := true
//	for flag{
//		//现在判定往哪个方向走,看上下左右四个方向中哪个方向的萝卜数最大,且萝卜数不为零
//		//建立四个result数组,决定往哪个方向走
//		result := make([]int, 4)
//		for i := 0; i < len(result); i++ {
//			//将上下左右四个元素的结果值存放在result切片当中
//			newM := road[i][0] + m //行
//			newN := road[i][1] + n //列
//			//判断是否满足条件
//			if newM >= 0 && newM < len(carrot) && newN >= 0 && newN < len(carrot[0]) {
//				result[i] = carrot[newM][newN]
//			}
//		}
//		//返回max和index值 确定它是往上/下/左/右走的,对m,n更新
//		max, index := maxNum(result)
//		if max != 0 {
//			//利用index值对m,n更新,前提是max值不为0
//			m += road[index][0]
//			n += road[index][1]
//			//加上最大值
//			carrotNum += carrot[m][n]
//			carrot[m][n] = 0
//		}else{
//			break
//		}
//	}
//	return carrotNum
//}
//func maxNum(num []int) (max int, index int) {
//	max = num[0]
//	for i := 1; i < len(num); i++ {
//		if num[i] > max {
//			max = num[i]
//			index = i
//		}
//	}
//	return max, index
//}

//func SquareArray(a []int) []int {
//	// write your code here
//	for i := 0; i < len(a); i++ {
//		a[i] = a[i] * a[i]
//	}
//	sort.Ints(a)
//	return a
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func CountNodes(head *ListNode) int {
//	// write your code here
//	lens := 0
//	for head != nil{
//		lens ++
//		head = head.Next
//	}
//	return lens
//}

//func SwapIntegers(a []int, index1 int, index2 int)  {
//	// write your code here
//	a[index1],a[index2] = a[index2],a[index1]
//}
//func Prime(n int) []int {
//	// write your code here
//	//isPrime()
//}

//func Prime(n int) []int {
//	// write your code here
//	slice := []int{}
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

//func Fibonacci(n int) int {
//	// write your code here
//	q, p := 0, 1
//	if n <= 1 {
//		return n - 1
//	}
//	for i := 2; i < n; i++ {
//		sum := q + p
//		q = p
//		p = sum
//	}
//	return p
//}

//func DigitCounts(k int, n int) int {
//	// write your code here
//}

//func MaxNum(nums []int) int {
//	// write your code here
//	sort.Ints(nums)
//	return nums[len(nums) - 1]
//}

//func LowercaseToUppercase2(letters string) string {
//	// write your code here
//	return strings.ToUpper(letters)
//}
//func MaxOfThreeNumbers(num1 int, num2 int, num3 int) int {
//	// write your code here
//	max1 := Max(num1, num2)
//	max2 := Max(max1, num3)
//	return max2
//}
//func Max(num1, num2 int) int {
//	max := 0
//	if num1 > num2 {
//		max = num1
//	} else {
//		max = num2
//	}
//	return max
//}

//func LowercaseToUppercase(character byte) byte {
//	// write your code here
//	character += 32
//	return character
//}
//func StrStr(source string, target string) int {
//	// Write your code here
//	//if target == ""{
//	//	return 0
//	//}
//	return strings.Index(source,target)
//}
//func ReverseInteger(number int) int {
//	// write your code here
//	i := number % 10
//	j := number / 10 % 10
//	k := number % 100
//	return k + j * 10 + i * 100
//}

/**
 * @param a: an integer array
 * @return: nothing
 */
//func SortIntegers(a []int)  {
//	// write your code here
//	sort.Ints(a)
//}
//func ReverseWords(s string) (res string) {
//	// write your code here
//	s = " " + s + " "
//	l, r := len(s) - 1, len(s) - 1
//	for i := len(s) - 2; i >= 0; i--{
//		if s[i] == ' '{
//			l, r = i, l
//			if r > l + 1{
//				res = res + s[l + 1:r] + " "
//			}
//		}
//	}
//	return res[:len(res) - 1]
//}

//func TwoSum(numbers []int, target int) []int {
//	// write your code here
//	for i := 0; i < len(numbers); i++ {
//		for j := 0; j < len(numbers); j++ {
//			if numbers[i]+numbers[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

//func AreaOfTrapezoid(a int, b int, h int) float64 {
//	// Write your code here
//	return float64(a + b) * float64(h) / 2.0
//}

//func GetSum(a int, b int) int {
//	// Write your code here
//	sum := 0
//	for i := a; i <= b; i++ {
//		if i%3 == 0 {
//			sum += i
//		}
//	}
//	return sum
//}

//func GetArea(r float64) float64 {
//	// write your code here
//	const pi = 3.14
//	return pi * r * r
//}

/**
 * @param str: A string
 * @return: a boolean
 */
//func IsUnique(str string) bool {
//	// write your code here
//	if len(str) == 0 {
//		return true
//	}
//	m := make(map[int]int)
//	for i := 0; i < len(str); i++ {
//		key := int(str[i])
//		for j := 0; j < len(m);j++ {
//			if m[j] == key {
//				return true
//			}
//		}
//		m[i] = key
//	}
//	return false
//}

//455.Assign Cookie(Easy)

//func findContentChildren(g []int, s []int) int {
//	//1.对数组 g s 进行排序,保证数组的有序性
//	sort.Ints(g)
//	sort.Ints(s)
//	//2.初始化孩子和饼干数
//	//child cookie 还代表着二者的index值(复用,避免多分配变量空间)
//	child, cookie := 0, 0
//	for child < len(g) && cookie < len(s) {
//		//若能被喂饱 则二者index均往后走一位
//		if g[child] <= s[cookie] {
//			child++
//			cookie++
//		} else {
//			//若当前饼干不足以喂饱最低饥饿度的孩子，则饼干索引值++
//			cookie++
//		}
//	}
//	return child
//}

//135. Candy(Hard)

//func candy(ratings []int) int {
//	size := len(ratings)
//	//branch 1: size < 2
//	if size < 2 {
//		return size
//	}
//	//branch 2: size >= 2
//	//1.初始化一个切片nums 存放每个孩子所需要的初始糖果数 1
//	nums := make([]int, size)
//	for i := 0; i < size; i++ {
//		nums[i] = 1
//	}
//	//2.左右两边遍历 --左遍历 右侧的值 > 左侧的值,candy数++
//	for i := 1; i < size; i++ {
//		if ratings[i] > ratings[i-1] {
//			nums[i] = nums[i-1] + 1
//		}
//	}
//	//右遍历 左侧的值 > 右侧且左侧的candy <= 右侧的candy,则左侧的candy为右侧的candy++。否侧保持不变
//	for i := size - 1; i > 0; i-- {
//		if ratings[i] < ratings[i-1] && nums[i-1] <= nums[i] {
//			nums[i-1] = nums[i] + 1
//		}
//	}
//	//3.对数组求和
//	sum := 0
//	for i := 0; i < size; i++ {
//		sum += nums[i]
//	}
//	return sum
//}

//435.Non-overlapping Intervals(Medium)

//func eraseOverlapIntervals(intervals [][]int) int {
//	//保留最小区间且不相交重叠的
//	//1.先把区间按结尾大小进行升序排序，每次选择结尾最小且和前一个选择的区间不重叠的区间。
//	size := len(intervals)
//	//1.判空
//	if size == 0 {
//		return 0
//	}
//	//2.利用sort包的SliceStable,保证稳定排序的前提下 将intervals[][]int中的items的结尾进行排序,
//	//小的放在前面，大的放在后面。由于intervals[]int内部size为2,则tail的index=1
//	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
//	//3.此时数组已经排序完成,初始化要移除的区间数量total 和前一个intervals[]中的tail元素
//	total, prev := 0, intervals[0][1]
//	//4.循环遍历数组切片区间
//	for i := 1; i < size; i++ {
//		if intervals[i][0] < prev {
//			total++
//		} else {
//			prev = intervals[i][1]
//		}
//	}
//	return total
//}

//605.Can Place Flowers
//func canPlaceFlowers(flowerbed []int, n int) bool {
//	//1.首尾两端添加0,避免边界条件的干扰
//	flowerbed = append([]int{0}, append(flowerbed, 0)...)
//	//2.遍历数组 i := 1 跳过index=0,i < len(flowerbed)-1跳过index=len-2
//	for i := 1; i < len(flowerbed)-1; i++ {
//		//3.判定当前位置是否栽种,若是则跳过当前位置
//		if flowerbed[i] == 1 {
//			i++
//			continue
//		}
//		//4.判断当前位置前后两个位置是否为未栽种,若是,则将当前位置置为1(表示栽种)
//		//i++ 往后轮询index值,n--代表可种花数减少一朵。
//		//此时已经得知当前位置为0(由经过上面的if-continue可知)
//		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
//			flowerbed[i] = 1
//			i++
//			n--
//		}
//		//5.走到这里所满足的条件是初始输入的n=0&&花坛中没有可以种植的位置
//		if n == 0 {
//			return true
//		}
//	}
//6.巧妙利用bool表达式
//True:
//走到此处的条件是初始输入的n>=0。在for中所种植的花朵数>=n。
//=0 即n==0, >n 即 n<0。
//比如:花坛能种2朵,n=1.经过for之后n=-1。
//False:
//初始的n>0。在for中并未种植一朵花。
//return n <= 0
//}
