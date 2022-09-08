package main

import "sort"

func StringPermutation2(str string) []string {
	// write your code here
	bytes := []byte(str)
	//按照字母序对str进行排序
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	visited := make([]bool, len(bytes))
	sub := make([]byte, 0, len(bytes))
	res := []string{}

	dfs(bytes, sub, visited, &res)

	return res
}

func dfs(bytes []byte, sub []byte, visited []bool, res *[]string) {
	if len(sub) == len(bytes) {
		array := make([]byte, len(bytes))
		copy(array, sub)
		*res = append(*res, string(array))
		return
	}

	for i, char := range bytes {
		if visited[i] {
			continue
		}
		if i > 0 && bytes[i-1] == bytes[i] && !visited[i-1] {
			continue
		}
		sub = append(sub, char)
		visited[i] = true
		dfs(bytes, sub, visited, res)
		sub = sub[:len(sub)-1]
		visited[i] = false
	}
}
