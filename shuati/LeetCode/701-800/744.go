package _01_800

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	//corner case
	if target >= letters[len(letters)-1] {
		return letters[0]
	}
	i := sort.Search(len(letters)-1, func(i int) bool { return letters[i] > target })
	return letters[i]
}
