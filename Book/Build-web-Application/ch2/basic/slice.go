package main

import "fmt"

func main() {
	// 和声明array一样，只是少了长度
	var fslice []int
	fmt.Println(fslice)
	//[]
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Printf("%c\n", slice)
	//fmt.Println(slice)
	//以 int 类型输出： [97 98 99 100]
	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(ar)
	//[97 98 99 100 101 102 103 104 105 106] ASCII 值
	// 声明两个含有byte的slice
	var a, b []byte
	fmt.Println(a, b)
	//[] []
	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	fmt.Println(a)
	//[99 100 101]
	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	fmt.Println(b)
	//[100 101]

	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(array)
	// 声明两个slice
	var aSlice, bSlice []byte
	fmt.Println("------", aSlice, bSlice)

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g
	fmt.Println(bSlice)

	var array1 [10]int
	slice1 := array[2:4]
	fmt.Println(array1)
	fmt.Println(len(slice1), cap(slice1)) //len = 2 cap = 8
	//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(slice1)
	//[99 100]
	slice1 = array[2:4:7]
	//切片可以包含两个或三个参数
	//array[2:4] 指定切片的起始范围 左取右不取 cap容量为原始数组-头(2) = 8
	//array[2:4:7] 指定切片的低值 高值 最大值 cap容量为 最大值-低值 = 5
	fmt.Println(slice1)
	fmt.Println(len(slice1), cap(slice1))

}
