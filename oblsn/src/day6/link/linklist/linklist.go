package linklist

import (
	"fmt"
)

type LinkNode struct {
	data interface{}
	next *LinkNode
}

type Link struct {
	head *LinkNode
	tail *LinkNode
}

func (p *Link) InsertHead(data interface{}) {
	node := &LinkNode{
		data: data,
		next: nil,
	}
	// 如果此链表中没有node 则新添加的这个node作为其首个node
	//防止链表中没有node，新添对象时空指针
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}
	node.next = p.head //新添加的节点先指向原node的头（地址）
	p.head = node      //头部新添了一个node，所以整个链表的的头是新添加的这个节点的地址
}

func (p *Link) InsertTail(data interface{}) {

	node := &LinkNode{
		data: data,
		next: nil,
	}
	if p.tail == nil && p.head == nil {
		p.tail = node
		p.head = node
		return
	}

	p.tail.next = node //链表中最后一个节点的next指针指向要添加的节点
	p.tail = node      //整个链表的尾是新添的这个节点的地址

}

//链表遍历
func (p *Link) Trans() {
	t := p.head
	for t != nil {
		fmt.Println(t.data)
		t = t.next
	}
}
