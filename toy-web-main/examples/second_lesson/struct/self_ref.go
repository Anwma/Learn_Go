package main

func main() {

}

type Node struct {
	//left Node
	//right Node

	//拓展点：压缩指针
	//自引用只能使用指针
	left *Node
	right *Node

	// 这个也会报错
	// nn NodeNode
}


type NodeNode struct {
	node Node
}