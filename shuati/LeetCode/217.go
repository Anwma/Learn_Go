package main

//排序后相邻元素不一
//func containsDuplicate(nums []int) bool {
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] == nums[i+1] {
//			return true
//		}
//	}
//	return false
//}

//hashmap

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, has := m[v]; has {
			return true
		}
		m[v] = i
	}
	return false
}
