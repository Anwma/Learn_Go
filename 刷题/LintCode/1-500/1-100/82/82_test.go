package main

// hashmap
func SingleNumber(a []int) int {
	// write your code here
	if a == nil || len(a) == 0 {
		return -1
	}
	m := make(map[int]int)
	for _, v := range a {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = v
		}
	}
	for _, v := range m {
		return v
	}
	return -1
}

// 位运算 异或
func SingleNumber(a []int) int {
	n := len(a)
	if a == nil || n == 0 {
		return -1
	}
	res := 0
	for i := 0; i < n; i++ {
		res ^= a[i]
	}
	return res
}
