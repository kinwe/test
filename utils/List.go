package utils

import "fmt"

type Node struct {
	Data interface{}   //定义数据域
	Next *Node //定义地址域(指向下一个表的地址)
}

type List struct {
	headNode *Node //头节点
}

func (this *List) IsEmpty() bool {

	if this.headNode == nil {
		return true
	} else {
		return false
	}
}

func (this *List) Length() int {
	//获取链表头结点
	cur := this.headNode
	//定义一个计数器，初始值为0
	count := 0

	for cur != nil {
		//如果头节点不为空，则count++
		count++
		//对地址进行逐个位移
		cur = cur.Next
	}
	return count
}

func (this *List) Add(data interface{}) *Node {
	node := &Node{Data: data}
	node.Next = this.headNode
	this.headNode = node
	return node
}

func (this *List) Append(data interface{}) {

	node := &Node{Data: data}
	if this.IsEmpty() {
		this.headNode = node
	} else {
		cur := this.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
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
