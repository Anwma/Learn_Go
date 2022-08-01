package main

import "sort"

type Interface interface {
	//sort 包排序提供了对切片和用户定义的集合进行排序的原语。
	sort.Interface      //嵌入字段sort.Interface
	Push(x interface{}) //a Push method to push elements into the heap
	Pop() interface{}   //a Pop elements that pops elements from the heap
}
type Interface1 interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less returns whether the element with index i should sort
	// before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

/*

// io.ReadWriter
type ReadWriter interface {
	Reader
	Writer
}

*/
