package main

import "sort"

// 暴力
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}

// 二分
func findTheDistanceValue1(arr1 []int, arr2 []int, d int) int {
	res := 0
	//排序
	sort.Ints(arr2)
	for _, v := range arr1 {
		//查找出现不满足
		ll := sort.Search(len(arr2), func(index int) bool {
			//arr2[index] >= v-d 规避掉了绝对值计算
			return arr2[index] >= v-d
		})
		hh := sort.Search(len(arr2), func(index int) bool {
			return arr2[index] > v+d
		})
		if ll == hh {
			res++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
