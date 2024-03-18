package linklist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head *ListNode
	len  uint
}

// NewListNode 新数据节点
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// GetNext 下一个节点
func (this *ListNode) GetNext() *ListNode {
	return this.next
}

// GetValue 获取值
func (this *ListNode) GetValue() interface{} {
	return this.value
}

// NewLinkedList 获取链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// InsertAfter 后面插入
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)
	newNode.next = p.next
	p.next = newNode
	this.len++

	return true
}

// InsertBefore 前面插入
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	//寻找head
	cur := this.head.next
	pre := this.head

	for nil != cur {
		if cur == p {
			break
		}

		pre = cur
		cur = pre.next
	}

	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	cur.next = newNode
	newNode.next = p
	this.len++
	return true
}

// InsertToHead 插入到开头
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// InsertToTail 插入到结尾
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}

	return this.InsertAfter(cur, v)
}

// FindByIndex 查找索引
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= this.len {
		return nil
	}

	cur := this.head.next

	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}

	return cur
}

// 删除节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}

	cur := this.head.next
	pre := this.head

	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if nil == cur {
		return false
	}

	pre.next = p.next
	p = nil
	this.len--
	return true

}

// 打印节点
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
