package dlist

import (
	"data_struct/utils"
)

/*

		   head
		 	|--------------------------- |
			v							 |
	    --Node1 <---> Node2 <---> Node3---
		|							^
		|---------------------------|
*/

type Node struct {
	prev  *Node
	next  *Node
	Value interface{}
	list  *List
}

func (n *Node) Prev() *Node {
	if n.list == nil {
		return nil
	}
	return n.prev
}

func (n *Node) Next() *Node {
	if n.list == nil {
		return nil
	}
	return n.next
}

type List struct {
	head *Node
	len  uint
}

func New() *List {
	list := &List{}
	return list
}

func (l *List) Len() uint {
	return l.len
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) PushBack(v interface{}) {
	node := &Node{Value:v}
	if l.head == nil {
		l.head = node
		node.prev = node
		node.next = node
	} else if l.len == 1 {
		l.head.prev = node
		l.head.next = node
		node.next = l.head
		node.prev = l.head
	} else {
		l.head.prev.next = node
		node.next = l.head
		node.prev = l.head.prev
		l.head.prev = node
	}
	node.list = l
	l.len++
}

func (l *List) PopBack() interface{} {
	ret := l.head.prev.Value
	l.Remove(int(l.Len()-1))
	return ret
}

func (l *List) PushFront(v interface{}) {
	node := &Node{Value:v}
	if l.head == nil {
		node.prev = node
		node.next = node
	} else if l.len == 1 {
		node.next = l.head
		node.prev = l.head
		l.head.prev = node
		l.head.next = node
	} else {
		l.head.prev.next = node
		node.next = l.head
		node.prev = l.head.prev
		l.head.prev = node

	}
	l.head = node
	node.list = l
	l.len++
}

func (l *List) PopFront() interface{} {
	ret := l.head.Value
	l.Remove(0)
	return ret
}

func (l *List) Insert(v interface{}, pos uint) {
	if pos > l.len {
		pos = l.len
	}
	node := &Node{Value:v}
	if l.head == nil {
		l.head = node
		node.prev = node
		node.next = node
	} else if l.len == 1 {
		node.next = l.head
		node.prev = l.head
		l.head.prev = node
		l.head.next = node
	} else {
		for cur := l.head; ; cur = cur.Next() {
			pos--
			if pos == 0 {
				node.next = cur.next
				cur.next.prev = node
				cur.next = node
				node.prev = cur
				break
			}
		}
	}
	node.list = l
	l.len++
}

func (l *List) Find(v interface{}) int {
	if l.head.Value == v {
		return 0
	}
	pos := 0
	for cur := l.head.Next(); cur != l.head; cur = cur.Next() {
		pos++
		if cur.Value == v {
			return pos
		}
	}
	return -1
}

func (l *List) Remove(pos int) {
	if pos < 0 || l.len == 0 || pos > int(l.len) {
		return
	}
	if pos == 0 {
		l.head.prev.next = l.head.next
		l.head.next.prev = l.head.prev
		l.head = l.head.next
	} else {
		for cur := l.head; ; cur = cur.Next() {
			if pos == 1 {
				del := cur.next
				del.next.prev = del.prev
				del.prev.next = del.next
				break
			}
			pos--
		}
	}
	l.len--
}

func (l *List) Empty() bool {
	return l.len == 0
}

func (l *List) Back() interface{} {
	return l.head.prev.Value
}

func (l *List) Front() interface{} {
	return l.head.Value
}

func (l *List) Traverse(visitor utils.Visitor) {
	for n := l.head; n != nil; n = n.Next() {
		visitor(n.Value)
	}
}
