package main

//暴力
//func NumJewelsInStones(j string, s string) int {
//	// Write your code here
//	count := 0
//	for _, v1 := range j {
//		for _, v2 := range s {
//			if v1 == v2{
//				count++
//			}
//		}
//	}
//	return count
//}

func NumJewelsInStones(j string, s string) int {
	// Write your code here
	m := make(map[rune]bool)
	count := 0
	for _, char := range j {
		m[char] = true
	}
	for _, char := range s {
		if _, ok := m[char]; ok {
			count++
		}
	}
	return count
}
