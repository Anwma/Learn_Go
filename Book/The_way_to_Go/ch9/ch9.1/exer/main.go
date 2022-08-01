package main

import (
	"fmt"
)

// DNode 节点
type DNode struct {
	data interface{}
	prev *DNode
	next *DNode
}

// DList 双向链表
type DList struct {
	size uint64
	head *DNode
	tail *DNode
}

func main() {
	//创建一个DList 链接
	myDlist := new(DList)
	fmt.Println(myDlist)
	myDlist.Append(101)
	myDlist.Append(102)
	myDlist.Append(103)
	myDlist.Append("Apple")
	myDlist.Append("pea")
	myDlist.Append("Banana")

	mynode := myDlist.getNode("pea")

	myDlist.InsertNext(mynode, "Orange")
	myDlist.InsertPrev(mynode, "watermelon")

	//var mynode *DNode = myDlist.getNode("pea")
	myDlist.DisplayData()
	fmt.Println()
	myDlist.Remove(mynode)
	myDlist.DisplayData()
}

// Init 初始化链表Init
func (dList *DList) Init() {
	_dlist := *(dList)
	_dlist.size = 0
	_dlist.head = nil
	_dlist.tail = nil
}

/*********************************
*    head       prev  tail
*    ********        *********
*    *       * <-----    *        *
*    *       * ----->    *        *
*    ********        *********
*              next
*********************************/

// Append 新增数据Append
func (dList *DList) Append(data interface{}) {
	newNode := new(DNode)  //new一个节点
	(*newNode).data = data //添加数据

	if (*dList).GetSize() == 0 { //初始化表头
		(*dList).head = newNode
		(*dList).tail = newNode
		(*newNode).prev = nil
		(*newNode).next = nil
	} else { //添加节点
		(*newNode).prev = (*dList).tail   //将新节点的prev指向表尾
		(*newNode).next = nil             //新节点next为空
		(*((*dList).tail)).next = newNode //将表尾的next指向newNode
		(*dList).tail = newNode           //将将表尾指向新节点
	}
	(*dList).size++ //容量++
}

// InsertNext 在节点后面插入数据
func (dList *DList) InsertNext(elmt *DNode, data interface{}) bool {
	if elmt == nil {
		return false
	}

	if dList.isTail(elmt) { //恰好在表尾
		dList.Append(data)
	} else {
		newNode := new(DNode)
		(*newNode).data = data
		(*newNode).prev = elmt         //(*newNode).prev指向elmt节点
		(*newNode).next = (*elmt).next //将(*newNode).next 指向节点(*elmt).next

		(*elmt).next = newNode            //将(*elmt).next指向newNode
		(*(*newNode).next).prev = newNode //将(*(*newNode).next).prev指向newNode节点
		(*dList).size++
	}
	return true
}

// InsertPrev 在节点之前插入数据
func (dList *DList) InsertPrev(elmt *DNode, data interface{}) bool {
	if elmt == nil {
		return false
	}

	if dList.isHead(elmt) { //新增加表头
		newNode := new(DNode)
		(*newNode).data = data
		(*newNode).next = dList.GetHead() //将(*newNode).next指向表头
		(*newNode).prev = nil             //将prev赋空

		(*(dList.head)).prev = newNode //将旧表头的prev指向newNode
		dList.head = newNode           //将newNode赋给表头
		dList.size++
		return true
	} else {
		prev := (*elmt).prev
		return dList.InsertNext(prev, data)
	}
}

// Remove 删除节点
func (dList *DList) Remove(elmt *DNode) interface{} {
	if elmt == nil {
		return 0
	}

	prev := (*elmt).prev
	next := (*elmt).next

	if dList.isHead(elmt) {
		dList.head = next
	} else {
		(*prev).next = next
	}

	if dList.isTail(elmt) {
		dList.tail = prev
	} else {
		(*next).prev = prev
	}

	dList.size--
	return (*elmt).GetData()
}

// MatchFun  匹配函数
func MatchFun(data1 interface{}, data2 interface{}) bool {
	if data1 != data2 {
		return false
	}
	return true
}

// Search 查找指定数据
func (dList *DList) Search(data interface{}) *DNode {
	if dList.GetSize() == 0 {
		return nil
	}

	node := dList.GetHead()
	for ; node != nil; node = node.GetNext() {
		if MatchFun(node.GetData(), data) == true {
			break
		}
	}
	return node
}

// DealData dealdata
func DealData(data interface{}) {
	switch data.(type) {
	case string:
		fmt.Printf("display all string data :%s\n", data)
	case int:
		fmt.Printf("display all int data : %d\n", data)
	default:
		fmt.Println("Unkown format type.")
	}
}

// DisplayData 显示数据
func (dList *DList) DisplayData() bool {
	if dList.GetSize() == 0 {
		return false
	}

	node := dList.GetHead()
	for ; node != nil; node = node.GetNext() {
		DealData(node.GetData())
	}

	return true
}

// GetSize 获取链表长度
func (dList *DList) GetSize() uint64 {
	return (*dList).size
}

//isTail 判断是否在表尾
func (dList *DList) isTail(elmt *DNode) bool {
	return dList.GetTail() == elmt
}

// GetTail 获取表尾节点
func (dList *DList) GetTail() *DNode {
	return (*dList).tail
}

//isHead 判断是否是头节点
func (dList *DList) isHead(elmt *DNode) bool {
	return dList.GetHead() == elmt
}

// GetHead 获取头节点
func (dList *DList) GetHead() *DNode {
	return (*dList).head
}

// GetData 获取节点内部数据
func (dNode *DNode) GetData() interface{} {
	return (*dNode).data
}

// GetNext 获取下一个节点
func (dNode *DNode) GetNext() *DNode {
	return (*dNode).next
}

// GetPrev 获取上一个节点
func (dNode *DNode) GetPrev() *DNode {
	return (*dNode).prev
}

//getNode 获取含有指定数据的第一个结点
func (dList *DList) getNode(e interface{}) *DNode {
	var p = dList.head
	for p != nil {
		//找到该数据所在结点
		if e == p.data {
			return p
		} else {
			p = p.next
		}
	}
	return nil
}
