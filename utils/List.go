package utils

import "fmt"

type Node struct {
	Data interface{} //定义数据域
	Next *Node       //定义地址域(指向下一个表的地址)
}

type List struct {
	count int64

	headNode *Node //头节点

	tailNode *Node
}

func (this *List) IsEmpty() bool {

	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

func (this *List) Length() int64 {

	return this.count
}

func (this *List) Add(data interface{}) {
	node := &Node{Data: data}

	if this.count > 0 {
		node.Next = this.headNode
		this.headNode = node
	} else {
		this.headNode = node
		this.tailNode = this.headNode
	}
	this.count++
}

func (this *List) Append(data interface{}) {

	node := &Node{Data: data}
	if this.count == 0 {
		this.headNode = node
		this.tailNode = this.headNode
		this.count++
	} else {
		this.tailNode.Next = node
		this.tailNode = this.tailNode.Next
		this.count++
	}
}

func (this *List) ReverseNode() *Node {
	//  先声明两个变量
	//  前一个节点
	var preNode *Node
	preNode = this.headNode
	//  后一个节点
	cur := new(Node)
	cur = this.headNode.Next
	preNode.Next = nil
	for cur != nil {
		nextNode := cur.Next
		cur.Next = preNode
		preNode = cur
		cur = nextNode
	}
	this.headNode = preNode
	return preNode
}

func (this *List) ShowList() {

	if !this.IsEmpty() {
		cur := this.headNode
		for {
			fmt.Println(cur.Data)
			if cur.Next != nil {
				cur = cur.Next
			} else {
				break
			}
		}
	}
}
